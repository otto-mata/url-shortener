package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"url-shortener/internal/config"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"
)

// NewRouter constructs the HTTP router
func NewRouter(cfg config.Config, store storage.Store) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	svc := shortener.NewService(cfg.BaseURL, store)
	h := NewHandler(svc)

	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", h.Shorten)
	})

	// Redirect route
	r.Get("/{code}", h.Resolve)

	return r
}
