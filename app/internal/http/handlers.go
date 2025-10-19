package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler groups HTTP handlers
 type Handler struct {
	Svc Shortener
 }

 // Shortener defines the operations needed from the shortener service
 type Shortener interface {
	Shorten(r *http.Request, url string) (string, error)
	Resolve(r *http.Request, code string) (string, error)
 }

func NewHandler(svc Shortener) *Handler {
	return &Handler{Svc: svc}
}

 type shortenRequest struct {
	URL string `json:"url"`
 }

 type shortenResponse struct {
	Code string `json:"code"`
	ShortURL string `json:"short_url"`
 }

 // Shorten handles POST /api/shorten
 func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	code, err := h.Svc.Shorten(r, req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := shortenResponse{Code: code, ShortURL: r.Host + "/" + code}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
 }

 // Resolve handles GET /{code}
 func (h *Handler) Resolve(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		http.NotFound(w, r)
		return
	}
	url, err := h.Svc.Resolve(r, code)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
 }
