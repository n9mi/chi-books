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

// update a book by id
func (r *BookRepositoryImpl) UpdateOne(id string, book *entity.Book) {
	// assume that the record is exists. double check that FindOne is called first
	bookPtr := getStorage().books[id]
	book.ID = bookPtr.ID
	book.CreatedAt = bookPtr.CreatedAt
	book.UpdatedAt = time.Now()

	getStorage().books[id] = book
}

// delete book by id
func (r *BookRepositoryImpl) DeleteOne(id string) {
	// assume that the record is exists. double check that FindOne is called first
	delete(getStorage().books, id)
}

// delete all k,v pair in books map
func (r *BookRepositoryImpl) Truncate() {
	for key := range getStorage().books {
		delete(getStorage().books, key)
	}
}
