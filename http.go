package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	handlers "github.com/tmthrgd/httphandlers"
	"go.tmthrgd.dev/insta.tmthrgd.dev/internal/assets"
	"go.tmthrgd.dev/insta.tmthrgd.dev/internal/templates"
	"go.tmthrgd.dev/vfshash"
)

const robots = "User-agent: *\nDisallow: /"

var assetNames = vfshash.NewAssetNames(assets.FileSystem)

var (
	errorTmpl = newTemplate("error.tmpl")
	indexTmpl = newTemplate("index.tmpl")
)

type errorData struct {
	StatusCode int
	Message    string
}

var notFoundData = &errorData{
	http.StatusNotFound,
	"The requested file was not found.",
}

// notFoundHandler returns a handler that serves a 404 error page.
func notFoundHandler() http.HandlerFunc {
	return handlers.Must(handlers.ServeErrorTemplate(http.StatusNotFound, errorTmpl, notFoundData, "text/html; charset=utf-8")).ServeHTTP
}

// methodNotAllowedHandler returns a handler that serves a 405 Method Not
// Allowed error.
func methodNotAllowedHandler() http.HandlerFunc {
	data := &errorData{
		http.StatusMethodNotAllowed,
		"The requested method is not allowed for this resource.",
	}
	h := handlers.Must(handlers.ServeErrorTemplate(http.StatusMethodNotAllowed, errorTmpl, data, "text/html; charset=utf-8"))
	return handlers.SetHeader(h, "Allow", "GET, HEAD").ServeHTTP
}

// indexHandler returns a handler that serves the index page.
func indexHandler() http.HandlerFunc {
	return handlers.Must(handlers.ServeTemplate("index.html", time.Now(), indexTmpl, nil)).ServeHTTP
}

// faviconHandler returns a handler that serves the favicon.ico file.
func faviconHandler() http.HandlerFunc {
	return http.FileServer(assetNames).ServeHTTP
}

// robotsHandler returns a handler that serves the robots.txt file.
func robotsHandler() http.HandlerFunc {
	return handlers.ServeString("robots.txt", time.Now(), robots).ServeHTTP
}

// assetsHandler returns a handler that serves site assets.
func assetsHandler() http.Handler {
	return http.StripPrefix("/assets", http.FileServer(filter.Skip(assets.FileSystem, excludeAssets)))
}

// errorHandler converts a handler with an error return to a http.HandlerFunc,
// sending a 500 Internal Server Error, or a 502 Bad Gateway where appropriate,
// to the client when an error is returned.
func errorHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		data := &errorData{
			StatusCode: http.StatusInternalServerError,
		}
		switch err := err.(type) {
		case httpError:
			switch err.StatusCode {
			case http.StatusNotFound:
				data = notFoundData
			default:
				data.StatusCode = http.StatusBadGateway
			}
		case *url.Error:
			// TODO: use errors.Is once go1.13 lands.
			if err.Err == errPrivateAccount {
				data.StatusCode = http.StatusForbidden
				data.Message = "This post belongs to a private Instagram account."
			}
		default:
			if os.IsNotExist(err) {
				data = notFoundData
			}
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(data.StatusCode)

		if data.Message == "" {
			data.Message = reflect.ValueOf(err).Type().String() + ": " + err.Error()
		}

		if templateExecute(w, errorTmpl, data) != nil {
			fmt.Fprintf(w, "%d %s: %s", data.StatusCode,
				http.StatusText(data.StatusCode),
				html.EscapeString(data.Message))
		}
	}
}

// templateExecute calls Execute on the given template, only writing out the result if
// execution was successful.
func templateExecute(w http.ResponseWriter, tmpl *template.Template, data interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	if n, err := buf.WriteTo(w); n == 0 {
		// Only forward the error if nothing was written.
		return err
	}

	return nil
}

// newTemplate parses the template at templates/name and returns a new
// html/template.Template. It panics if the source is invalid or the template
// doesn't exist.
func newTemplate(name string) *template.Template {
	t := template.New(name).Funcs(templateFuncs)
	return template.Must(vfstemplate.ParseFiles(templates.FileSystem, t, name))
}

var templateFuncs = template.FuncMap{
	"httpStatusText": http.StatusText,
	"assetPath":      assetPath,
}

// assetPath returns the path to a named asset file.
func assetPath(name string) string {
	return path.Join("/assets/", assetNames.Lookup(name))
}

// excludeAssets returns true if the file should be excluded from the assets handler.
func excludeAssets(path string, info os.FileInfo) bool {
	return info.IsDir() || strings.HasPrefix(info.Name(), ".")
}

type httpError struct{ *http.Response }

func (he httpError) Error() string {
	return "upstream HTTP error: " + he.Response.Request.URL.String() + ": " + he.Response.Status
}
