package router

import (
	"chi-books/handler"

	"github.com/go-chi/chi/v5"
)

func BookRouter(r chi.Router, bookHandler *handler.BookHandler) {
	r.Route("/books", func(r chi.Router) {
		r.Get("/", bookHandler.GetAll)
		r.Post("/", bookHandler.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", bookHandler.GetByID)
			r.Put("/", bookHandler.Update)
			r.Delete("/", bookHandler.Delete)
		})
	})
}
