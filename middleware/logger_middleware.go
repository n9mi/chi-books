package middleware

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

// logs method, path, response time, etc
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		httplog.LogEntrySetField(ctx, "user", slog.StringValue(r.UserAgent()))

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
