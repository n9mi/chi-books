package main

import (
	"chi-books/config"
	"net/http"
)

func main() {
	logger := config.NewLogger()
	app := config.NewApp(logger)

	logger.Logger.Info("Server listening at :3000")
	http.ListenAndServe(":3000", app)
}
