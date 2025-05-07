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

// costumized JSON response writer for single message error response
func ErrorJSONResponse(w http.ResponseWriter, httpErr *HTTPError) {
	errRes := ErrorResponse{
		StatusCode: httpErr.StatusCode,
		Success:    false,
		Errors: map[string]string{
			"_error": httpErr.Error(),
		}}

	response.JSONResponse(w, errRes.StatusCode, errRes)
}

// costumized JSON response writer for structured error response (ex: request validation)
func ErrorWithFieldsJSONResponse(w http.ResponseWriter, statusCode int, fields map[string]string) {
	errRes := ErrorResponse{
		StatusCode: statusCode,
		Success:    false,
		Errors:     fields,
	}

	response.JSONResponse(w, errRes.StatusCode, errRes)
}
