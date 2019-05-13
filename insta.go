package main // import "tmthrgd.dev/go/insta.tmthrgd.dev"

//go:generate go run -tags=dev internal/assets/generate.go

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tmthrgd/gziphandler"
	handlers "github.com/tmthrgd/httphandlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(
		middleware.GetHead,
		handlers.SecurityHeadersWrap(&handlers.SecurityHeaders{
			ContentSecurityPolicy: "default-src 'none'; script-src 'self'; style-src 'self' https:; img-src 'self' https:; font-src https:; frame-ancestors 'none'; block-all-mixed-content; sandbox allow-forms allow-scripts; base-uri 'none'; report-uri https://tomthorogood.report-uri.com/r/d/csp/enforce",
		}),
		handlers.SetHeaderWrap("Server", "insta.tmthrgd.dev"),
		gziphandler.Wrapper(),
	)
	router.NotFound(notFoundHandler())

	// Asset routes
	router.Group(func(r chi.Router) {
		rr := r.With(
			handlers.SetHeaderWrap("Cache-Control", "public, max-age=1209600"), // 14 days
		)
		rr.Get("/favicon.ico", faviconHandler())
		rr.Get("/robots.txt", robotsHandler())

		r.With(
			handlers.NeverModified,
			handlers.SetHeaderWrap("Cache-Control", "public, max-age=31536000, immutable"), // 1 year
		).Mount("/assets", assetsHandler())
	})

	// HTML page routes
	router.Group(func(r chi.Router) {
		r.With(
			handlers.SetHeaderWrap("Cache-Control", "public, max-age=3600"), // 1 hour
		).Get("/", indexHandler())

		r.With(middleware.NoCache).Get("/p", pRedirectHandler())

		post := postHandler()
		rr := r.With(
			handlers.SetHeaderWrap("Cache-Control", "public, max-age=300"), // 5 minutes
		)
		rr.Get("/p/{postID}/", post)
		rr.Get("/p/{postID}/json", post)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
