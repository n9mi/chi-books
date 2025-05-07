package middleware

import (
	"chi-books/exception"
	"chi-books/util"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func CustomErrorMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if errConv, ok := err.(validator.ValidationErrors); ok {
					errFields := make(map[string]string)

					for _, e := range errConv {
						eFieldLower := util.CamelToSnakeCase(e.Field())

						switch e.Tag() {
						case "required":
							errFields[eFieldLower] = fmt.Sprintf("%s is required", eFieldLower)
						case "min":
							errFields[eFieldLower] = fmt.Sprintf("%s is should be more or equal than %s character", eFieldLower, e.Param())
						case "max":
							errFields[eFieldLower] = fmt.Sprintf("%s is should be more or equal than %s character", eFieldLower, e.Param())
						case "gte":
							errFields[eFieldLower] = fmt.Sprintf("%s is should be greater or equal than %s", eFieldLower, e.Param())
						case "lte":
							errFields[eFieldLower] = fmt.Sprintf("%s is should be less or equal than %s", eFieldLower, e.Param())
						}
					}

					exception.ErrorWithFieldsJSONResponse(w, http.StatusBadRequest, errFields)
				} else if errConv, ok := err.(*exception.HTTPError); ok {
					exception.ErrorJSONResponse(w, errConv)
				} else {
					defaultErr := &exception.HTTPError{
						StatusCode: http.StatusInternalServerError,
						Message:    "something wrong",
					}

					exception.ErrorJSONResponse(w, defaultErr)
				}
			}
		}()

		h.ServeHTTP(w, r)
	})
}
