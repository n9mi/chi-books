package service

import (
	"chi-books/dto/response"
	"chi-books/repository"
	"context"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	Validator      *validator.Validate
}

func NewBookService(
	bookRepository repository.BookRepository,
	validate *validator.Validate) BookService {

	return &BookServiceImpl{
		BookRepository: bookRepository,
		Validator:      validate,
	}
}

// find all books
func (s *BookServiceImpl) FindAll(ctx context.Context) []*response.BookResponse {
	booksE := s.BookRepository.FindAll()

	booksRes := make([]*response.BookResponse, len(booksE))
	for i, bE := range booksE {
		booksRes[i] = response.ToBookResponse(bE)
	}

	return booksRes
}
