package routers

import (
	handlers "github.com/B-AJ-Amar/go-bookstore-demo/bookstore/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthorRoutes(r chi.Router) {
	r.Route("/authors", func(r chi.Router) {
		r.Get("/", handlers.GetAuthors)
		r.Post("/", handlers.CreateAuthor)
		// r.Route("/{id:[0-9]+}", func(r chi.Router) {
		// 	r.Get("/", handlers.GetAuthor)
		// 	r.Put("/", handlers.UpdateAuthor)
		// 	r.Delete("/", handlers.DeleteAuthor)
		// })
	})
}