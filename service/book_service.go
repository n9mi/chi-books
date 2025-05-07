package service

import (
	"chi-books/dto/request"
	"chi-books/dto/response"
	"context"
)

type BookService interface {
	FindAll(ctx context.Context) []*response.BookResponse
	Create(ctx context.Context, req *request.BookRequest) (*response.BookResponse, error)
}
