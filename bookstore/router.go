package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger) // Log the HTTP requests
	r.Use(middleware.Recoverer) // Recover from panics
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Bookstore!"))
	})
	return r
}