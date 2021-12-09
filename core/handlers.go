package core

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health Check"))
	})

	return r
}
