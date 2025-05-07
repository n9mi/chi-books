package handler

import (
	"chi-books/dto/response"
	"chi-books/service"
	"net/http"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		BookService: bookService,
	}
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := h.BookService.FindAll(r.Context())

	response.DataJSONResponse(w, http.StatusOK, true, res)
}
