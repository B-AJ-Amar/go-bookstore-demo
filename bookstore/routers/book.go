package routers

import (
	handler "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/handler"
	"github.com/go-chi/chi/v5"
)

func BookRoutes(r chi.Router) {
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetBooks)
		r.Post("/", handler.CreateBook)
		r.Route("/{id:[0-9]+}", func(r chi.Router) {
			r.Get("/", handler.GetBook)
			r.Put("/", handler.UpdateBook)
			r.Delete("/", handler.DeleteBook)
		})
	})
}