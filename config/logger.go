package config

import (
	"log/slog"

	"github.com/go-chi/httplog/v2"
)

// get debug level logger instance
func NewLogger() *httplog.Logger {
	return httplog.NewLogger("chi-books", httplog.Options{
		LogLevel: slog.LevelDebug,
	})
}
