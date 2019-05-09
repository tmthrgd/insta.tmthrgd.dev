package main

//go:generate go run -tags=dev internal/assets/generate.go

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "github.com/tmthrgd/httphandlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(
		middleware.GetHead,
		handlers.SecurityHeadersWrap(nil),
		handlers.SetHeaderWrap("Server", "insta.tmthrgd.dev"),
	)
	router.NotFound(notFoundHandler())

	// Asset routes
	router.Group(func(r chi.Router) {
		r.Get("/favicon.ico", faviconHandler())
		r.Get("/robots.txt", robotsHandler())

		r.Mount("/assets", assetsHandler())
	})

	// HTML page routes
	router.Group(func(r chi.Router) {
		r.Get("/", indexHandler())
		r.Get("/p/{postID}/", postHandler())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
