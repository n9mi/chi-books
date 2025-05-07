package exception

import (
	"net/http"
)

// default not found handler
func CustomNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	errRes := &HTTPError{
		StatusCode: http.StatusNotFound,
		Message:    "not found resource",
	}

	ErrorJSONResponse(w, errRes)
}

// default method not allowed handler
func CustomMethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	errRes := &HTTPError{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "method not allowed",
	}

	ErrorJSONResponse(w, errRes)
}
