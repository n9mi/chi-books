package repository

import "chi-books/entity"

type BookRepository interface {
	FindOne(id string) (*entity.Book, bool)
	FindAll() []*entity.Book
	InsertOne(book *entity.Book)
	UpdateOne(id string, book *entity.Book)
	DeleteOne(id string)
	Truncate()
}
