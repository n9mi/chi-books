package test

import (
	"chi-books/config"
	"chi-books/repository"

	"github.com/go-chi/chi/v5"
)

const bookURL = "/api/v1/books"

var app *chi.Mux
var bookRepository repository.BookRepository

func init() {
	logger := config.NewLogger()

	app = config.NewApp(logger)
	bookRepository = repository.NewBookRepository()
}
