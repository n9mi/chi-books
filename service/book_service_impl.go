package service

import (
	"chi-books/dto/request"
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

// create a book
func (s *BookServiceImpl) Create(ctx context.Context, req *request.BookRequest) (*response.BookResponse, error) {
	if err := s.Validator.Struct(req); err != nil {
		return nil, err
	}

	bookE := req.ToEntity()
	s.BookRepository.InsertOne(bookE)

	bookRes := response.ToBookResponse(bookE)
	return bookRes, nil
}
