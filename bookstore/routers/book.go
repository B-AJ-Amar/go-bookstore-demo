package routers

import (
	handlers "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/handlers"
	"github.com/go-chi/chi/v5"
)

func BookRoutes(r chi.Router) {
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handlers.GetBooks)
		r.Post("/", handlers.CreateBook)
		r.Route("/{id:[0-9]+}", func(r chi.Router) {
			r.Get("/", handlers.GetBook)
			r.Put("/", handlers.UpdateBook)
			r.Delete("/", handlers.DeleteBook)
		})
	})
}