package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	app := chi.NewRouter()

	return app
}
