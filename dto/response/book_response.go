package response

import (
	"chi-books/entity"
	"time"
)

type BookResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedYear int       `json:"published_year"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToBookResponse(e *entity.Book) *BookResponse {
	return &BookResponse{
		ID:            e.ID,
		Title:         e.Title,
		Author:        e.Author,
		PublishedYear: e.PublishedYear,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}
