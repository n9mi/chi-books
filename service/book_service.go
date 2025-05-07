package service

import (
	"chi-books/dto/response"
	"context"
)

type BookService interface {
	FindAll(ctx context.Context) []*response.BookResponse
}
