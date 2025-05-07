package exception

import (
	"chi-books/dto/response"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int               `json:"status_code"`
	Success    bool              `json:"success"`
	Errors     map[string]string `json:"errors"`
}

func ErrorJSONResponse(w http.ResponseWriter, httpErr *HTTPError) {
	errRes := ErrorResponse{
		StatusCode: httpErr.StatusCode,
		Success:    false,
		Errors: map[string]string{
			"_error": httpErr.Error(),
		}}

	response.JSONResponse(w, errRes.StatusCode, errRes)
}

func ErrorWithFieldsJSONResponse(w http.ResponseWriter, statusCode int, fields map[string]string) {
	errRes := ErrorResponse{
		StatusCode: statusCode,
		Success:    false,
		Errors:     fields,
	}

	response.JSONResponse(w, errRes.StatusCode, errRes)
}
