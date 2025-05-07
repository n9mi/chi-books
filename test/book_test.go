package test

import (
	"chi-books/dto/response"
	"chi-books/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllBooks(t *testing.T) {
	type GetAllBooksResponse struct {
		StatusCode int                      `json:"status_code"`
		Success    bool                     `json:"success"`
		Data       []*response.BookResponse `json:"data"`
	}

	numBooks := 10
	newBooks := make([]entity.Book, numBooks)
	for i := 0; i < numBooks; i++ {
		newBooks[i] = entity.Book{
			Title:         fmt.Sprintf("Title %d", i+1),
			Author:        fmt.Sprintf("Author %d", i+1),
			PublishedYear: 2000 + i,
		}
		bookRepository.InsertOne(&newBooks[i])
	}

	t.Run("Success_200", func(t *testing.T) {
		req := newRequest(http.MethodGet, bookURL, nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)

		var res GetAllBooksResponse
		err := json.NewDecoder(rec.Body).Decode(&res)
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.True(t, res.Success)
		require.Equal(t, numBooks, len(res.Data))
	})

	bookRepository.Truncate()
}

func TestCreateBook(t *testing.T) {
	type CreateBookSuccessResponse struct {
		StatusCode int                    `json:"status_code"`
		Success    bool                   `json:"success"`
		Data       *response.BookResponse `json:"data"`
	}

	t.Run("Success_200", func(t *testing.T) {
		reqBody := map[string]interface{}{
			"title":          "coba title",
			"author":         "coba author",
			"published_year": 2020,
		}
		req := newRequest(http.MethodPost, bookURL, reqBody)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)

		var res CreateBookSuccessResponse
		err := json.NewDecoder(rec.Body).Decode(&res)
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.True(t, res.Success)
		require.NotEmpty(t, res.Data.ID)
		require.NotEmpty(t, res.Data.CreatedAt)
		require.NotEmpty(t, res.Data.UpdatedAt)
		require.Equal(t, reqBody["title"], res.Data.Title)
		require.Equal(t, reqBody["author"], res.Data.Author)
		require.Equal(t, reqBody["published_year"], res.Data.PublishedYear)
	})

	bookRepository.Truncate()
}
