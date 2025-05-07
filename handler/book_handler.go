package handler

import (
	"chi-books/dto/request"
	"chi-books/dto/response"
	"chi-books/exception"
	"chi-books/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		BookService: bookService,
	}
}

// get all books in form of slice
func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := h.BookService.FindAll(r.Context())

	response.DataJSONResponse(w, http.StatusOK, true, res)
}

// get single book by id
func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.BookService.FindOne(r.Context(), id)
	if err != nil {
		panic(err)
	}

	response.DataJSONResponse(w, http.StatusOK, true, res)
}

// create a new book
func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := new(request.BookRequest)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errConv := &exception.HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to parse JSON",
		}
		panic(errConv)
	}

	res, err := h.BookService.Create(r.Context(), req)
	if err != nil {
		panic(err)
	}

	response.DataJSONResponse(w, http.StatusOK, true, res)
}

// update existing book. returns error when book id doesn't exists
func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := new(request.BookRequest)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errConv := &exception.HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to parse JSON",
		}
		panic(errConv)
	}

	res, err := h.BookService.Update(r.Context(), id, req)
	if err != nil {
		panic(err)
	}

	response.DataJSONResponse(w, http.StatusOK, true, res)
}

// delete existing book. returns error when book id doesn't exists
func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.BookService.Delete(r.Context(), id)
	if err != nil {
		panic(err)
	}

	response.DataJSONResponse(w, http.StatusOK, true, nil)
}
