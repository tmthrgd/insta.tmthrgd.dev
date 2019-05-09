package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"reflect"
	"time"

	handlers "github.com/tmthrgd/httphandlers"
	"tmthrgd.dev/go/insta.tmthrgd.dev/internal/assets"
)

const error404 = `<!doctype html>
<meta charset=utf-8>
<meta name=viewport content="width=device-width,initial-scale=1">
<title>404 Not Found – insta.tmthrgd.dev</title>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css integrity="sha256-l85OmPOjvil/SOvVt3HnSSjzF1TUMyT9eV0c2BzEGzU=" crossorigin=anonymous>
<link rel=stylesheet href=/assets/style.css>
<h1>404 Not Found</h1>
<p>The requested file was not found.</p>`

var error500 = newTemplate(`<!doctype html>
<meta charset=utf-8>
<meta name=viewport content="width=device-width,initial-scale=1">
<title>500 Internal Server Error – insta.tmthrgd.dev</title>
<link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css integrity="sha256-l85OmPOjvil/SOvVt3HnSSjzF1TUMyT9eV0c2BzEGzU=" crossorigin=anonymous>
<link rel=stylesheet href=/assets/style.css>
<h1>500 Internal Server Error</h1>
<p>{{.Type}}: {{.Message}}</p>`)

const robots = "User-agent: *\nDisallow: /"

// notFoundHandler returns a handler that serves a 404 error page.
func notFoundHandler() http.HandlerFunc {
	return handlers.ServeError(http.StatusNotFound, []byte(error404), "text/html; charset=utf-8").ServeHTTP
}

// faviconHandler returns a handler that serves the favicon.ico file.
func faviconHandler() http.HandlerFunc {
	return http.FileServer(&noDirFileSystem{assets.FileSystem}).ServeHTTP
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

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)

		error500.Execute(w, &struct {
			Type, Message string
		}{
			reflect.ValueOf(err).Type().String(),
			err.Error(),
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
	return template.Must(template.New("").Parse(source))
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
