package config

import (
	"chi-books/handler"
	"chi-books/repository"
	"chi-books/router"
	"chi-books/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func NewApp(logger *httplog.Logger) *chi.Mux {
	r := chi.NewRouter()

	// setup dependencies
	validator := NewValidator()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, validator)
	bookHandler := handler.NewBookHandler(bookService)

	// register handler
	r.Route("/api/v1", func(r chi.Router) {
		router.BookRouter(r, bookHandler)
	})

	return r
}
