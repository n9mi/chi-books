package config

import (
	"chi-books/exception"
	"chi-books/handler"
	"chi-books/middleware"
	"chi-books/repository"
	"chi-books/router"
	"chi-books/service"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func NewApp(logger *httplog.Logger) *chi.Mux {
	r := chi.NewRouter()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	// setup dependencies
	validator := NewValidator()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, validator)
	bookHandler := handler.NewBookHandler(bookService)

	// setup middleware
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Logger)
	r.Use(middleware.CustomErrorMiddleware)

	// setup custom handler
	r.NotFound(exception.CustomNotFoundHandler)
	r.MethodNotAllowed(exception.CustomMethodNotAllowedHandler)

	// register handler
	r.Route("/api/v1", func(r chi.Router) {
		router.BookRouter(r, bookHandler)
	})

	return r
}
