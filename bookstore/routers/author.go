package routers

import (
	handler "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/handler"
	"github.com/go-chi/chi/v5"
)

func AuthorRoutes(r chi.Router) {
	r.Route("/authors", func(r chi.Router) {
		r.Get("/", handler.GetAuthors)
		r.Post("/", handler.CreateAuthor)
		r.Route("/{id:[0-9]+}", func(r chi.Router) {
			r.Get("/", handler.GetAuthor)
			r.Put("/", handler.UpdateAuthor)
			r.Delete("/", handler.DeleteAuthor)
		})
	})
}