package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"url-shortener/internal/storage"
)

// Service provides URL shortening logic
 type Service struct {
	baseURL string
	store   storage.Store
 }

 func NewService(baseURL string, store storage.Store) *Service {
	return &Service{baseURL: strings.TrimRight(baseURL, "/"), store: store}
 }

 // Shorten creates a short code for the given URL and stores the mapping
 func (s *Service) Shorten(_ *http.Request, url string) (string, error) {
	if url == "" {
		return "", errors.New("empty url")
	}
	code := generateCode(6)
	if err := s.store.Save(code, url); err != nil {
		return "", err
	}
	return code, nil
 }

 // Resolve returns the original URL for the given code
 func (s *Service) Resolve(_ *http.Request, code string) (string, error) {
	if url, ok := s.store.Get(code); ok {
		return url, nil
	}
	return "", errors.New("not found")
 }

 func generateCode(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	// URL-safe base64, remove non-alphanum for simplicity
	code := base64.RawURLEncoding.EncodeToString(b)
	if len(code) > n {
		code = code[:n]
	}
	return code
 }
