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

type booksResponse struct {
	StatusCode int                      `json:"status_code"`
	Success    bool                     `json:"success"`
	Data       []*response.BookResponse `json:"data"`
}

type bookResponse struct {
	StatusCode int                    `json:"status_code"`
	Success    bool                   `json:"success"`
	Data       *response.BookResponse `json:"data"`
}

func TestGetAllBooks(t *testing.T) {
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

		var res booksResponse
		err := json.NewDecoder(rec.Body).Decode(&res)
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.True(t, res.Success)
		require.Equal(t, numBooks, len(res.Data))
	})

	bookRepository.Truncate()
}

func TestGetByIDBook(t *testing.T) {
	newBook := entity.Book{
		Title:         "Random Title",
		Author:        "Random Author",
		PublishedYear: 2013,
	}
	bookRepository.InsertOne(&newBook)

	t.Run("Success_200", func(t *testing.T) {
		bookURLWithID := bookURL + "/" + newBook.ID
		req := newRequest(http.MethodGet, bookURLWithID, nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)

		var res bookResponse
		err := json.NewDecoder(rec.Body).Decode(&res)
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.True(t, res.Success)
		require.Equal(t, newBook.ID, res.Data.ID)
		require.Equal(t, newBook.Title, res.Data.Title)
		require.Equal(t, newBook.Author, res.Data.Author)
		require.Equal(t, newBook.PublishedYear, res.Data.PublishedYear)
	})

	bookRepository.Truncate()
}

func TestCreateBook(t *testing.T) {
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

		var res bookResponse
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

func TestUpdateBook(t *testing.T) {
	newBook := entity.Book{
		Title:         "Random Title",
		Author:        "Random Author",
		PublishedYear: 2013,
	}
	bookRepository.InsertOne(&newBook)

	t.Run("Success_200", func(t *testing.T) {
		bookURLWithID := bookURL + "/" + newBook.ID
		reqBody := map[string]interface{}{
			"title":          "Edit Title",
			"author":         "Edit Author",
			"published_year": 2023,
		}
		req := newRequest(http.MethodPut, bookURLWithID, reqBody)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)

		var res bookResponse
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

func TestDeleteBook(t *testing.T) {
	newBook := entity.Book{
		Title:         "Random Title",
		Author:        "Random Author",
		PublishedYear: 2013,
	}
	bookRepository.InsertOne(&newBook)

	t.Run("Success_200", func(t *testing.T) {
		bookURLWithID := bookURL + "/" + newBook.ID
		req := newRequest(http.MethodDelete, bookURLWithID, nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code)

		var res bookResponse
		err := json.NewDecoder(rec.Body).Decode(&res)
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		require.True(t, res.Success)
		require.Nil(t, res.Data)
	})

	bookRepository.Truncate()
}
