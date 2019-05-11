package main

import (
	"bytes"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"time"

	handlers "github.com/tmthrgd/httphandlers"
	"tmthrgd.dev/go/insta.tmthrgd.dev/internal/assets"
	"tmthrgd.dev/go/vfshash"
)

var errorTmpl = newTemplate(`<!doctype html>
<meta charset=utf-8>
<meta name=viewport content="width=device-width,initial-scale=1">
<title>{{.StatusCode}} {{httpStatusText .StatusCode}} – insta.tmthrgd.dev</title>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css integrity="sha256-l85OmPOjvil/SOvVt3HnSSjzF1TUMyT9eV0c2BzEGzU=" crossorigin=anonymous>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css integrity="sha256-2YQRJMXD7pIAPHiXr0s+vlRWA7GYJEK0ARns7k2sbHY=" crossorigin=anonymous>
<link rel=stylesheet href="https://fonts.googleapis.com/css?family=Raleway">
<link rel=stylesheet href={{assetPath "/error.css"}}>
<main class=container>
<h1>{{.StatusCode}} {{httpStatusText .StatusCode}}</h1>
<p>{{.Message}}</p>
</main>`)

const robots = "User-agent: *\nDisallow: /"

var assetNames = vfshash.NewAssetNames(assets.FileSystem)

type errorData struct {
	StatusCode int
	Message    string
}

// notFoundHandler returns a handler that serves a 404 error page.
func notFoundHandler() http.HandlerFunc {
	h, err := handlers.ServeErrorTemplate(http.StatusNotFound, errorTmpl, &errorData{
		http.StatusNotFound,
		"The requested file was not found.",
	}, "text/html; charset=utf-8")
	if err != nil {
		panic(err)
	}

	return h.ServeHTTP
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
	return http.StripPrefix("/assets", http.FileServer(&noDirFileSystem{assets.FileSystem}))
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

		var (
			statusCode = http.StatusInternalServerError
			msg        string
		)
		switch err := err.(type) {
		case httpError:
			statusCode = http.StatusBadGateway

			if err.StatusCode == http.StatusNotFound {
				statusCode = http.StatusNotFound
				msg = "The requested file was not found."
			}
		case *url.Error:
			// TODO: use errors.Is once go1.13 lands.
			if err.Err == errPrivateAccount {
				statusCode = http.StatusForbidden
				msg = "This post belongs to a private Instagram account."
			}
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(statusCode)

		if msg == "" {
			msg = reflect.ValueOf(err).Type().String() + ": " + err.Error()
		}

		errorTmpl.Execute(w, &errorData{
			statusCode,
			msg,
		})
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

// newTemplate parses source and returns a new html/template.Template. It
// panics if source is invalid.
func newTemplate(source string) *template.Template {
	return template.Must(template.New("").Funcs(templateFuncs).Parse(source))
}

var templateFuncs = template.FuncMap{
	"httpStatusText": http.StatusText,
	"assetPath":      assetPath,
}

// assetPath returns the path to a named asset file.
func assetPath(name string) string {
	return path.Join("/assets/", assetNames.Lookup(name))
}

type noDirFileSystem struct{ http.FileSystem }

func (fs *noDirFileSystem) Open(name string) (http.File, error) {
	f, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}

	if stat, err := f.Stat(); err == nil && stat.IsDir() {
		f.Close()
		return nil, os.ErrNotExist
	}

	return f, nil
}

type httpError struct{ *http.Response }

func (he httpError) Error() string {
	return "upstream HTTP error: " + he.Response.Request.URL.String() + ": " + he.Response.Status
}
