package response

import (
	"encoding/json"
	"net/http"
)

// JSON response writer
func JSONResponse(w http.ResponseWriter, statusCode int, content interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(content)
}

// costumized JSON response writer for structured data response
func DataJSONResponse(w http.ResponseWriter, statusCode int, success bool, data interface{}) {
	res := DataResponse{
		StatusCode: statusCode,
		Success:    success,
		Data:       data,
	}

	JSONResponse(w, statusCode, res)
}
