package request

import "chi-books/entity"

type BookRequest struct {
	Title         string `json:"title" validate:"required,min=1,max=100"`
	Author        string `json:"author" validate:"required,min=1,max=100"`
	PublishedYear int    `json:"published_year" validate:"required,gte=1"`
}

func (r *BookRequest) ToEntity() *entity.Book {
	return &entity.Book{
		Title:         r.Title,
		Author:        r.Author,
		PublishedYear: r.PublishedYear,
	}
}
