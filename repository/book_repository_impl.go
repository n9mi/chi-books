package repository

import (
	"chi-books/entity"
	"time"

	"github.com/google/uuid"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

// find a book by id
func (r *BookRepositoryImpl) FindOne(id string) (*entity.Book, bool) {
	book, exists := getStorage().books[id]

	return book, exists
}

// get all books as slice
func (r *BookRepositoryImpl) FindAll() []*entity.Book {
	books := make([]*entity.Book, len(getStorage().books))

	ctr := 0
	for _, book := range getStorage().books {
		books[ctr] = book
		ctr++
	}

	return books
}

// insert book entity to map with {key: string = book.ID, value: *entity.Book = *book} pair
func (r *BookRepositoryImpl) InsertOne(book *entity.Book) {
	book.ID = uuid.NewString()
	timeNow := time.Now()
	book.CreatedAt = timeNow
	book.UpdatedAt = timeNow

	getStorage().books[book.ID] = book
}
