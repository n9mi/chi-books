package service

import (
	"chi-books/dto/request"
	"chi-books/dto/response"
	"chi-books/exception"
	"context"
)

type BookService interface {
	FindAll(ctx context.Context) []*response.BookResponse
	FindOne(ctx context.Context, id string) (*response.BookResponse, *exception.HTTPError)
	Create(ctx context.Context, req *request.BookRequest) (*response.BookResponse, error)
	Update(ctx context.Context, id string, req *request.BookRequest) (*response.BookResponse, error)
	Delete(ctx context.Context, id string) error
}
