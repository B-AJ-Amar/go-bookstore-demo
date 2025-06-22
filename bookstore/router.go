package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	routers "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/routers"
	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/utils"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(func(h http.Handler) http.Handler {
		return otelhttp.NewHandler(h, "request")
	})

	r.Use(middleware.Logger) // Log the HTTP requests
	r.Use(middleware.Recoverer) // Recover from panics
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Bookstore!"))
	})

	r.Route("/api", func(r chi.Router) {

		r.Use(middleware.AllowContentType("application/json")) // Allow only JSON content type
		r.Use(utils.JSONContentTypeMiddleware)

		routers.BookRoutes(r)
		routers.AuthorRoutes(r)

	})

	return r
}