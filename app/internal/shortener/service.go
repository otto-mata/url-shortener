package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"url-shortener/internal/storage"
)

// Service provides URL shortening logic
type Service struct {
	baseURL string
	store   storage.Store
}

type linkStats struct {
	Code   string `json:"code"`
	Target string `json:"target_url"`
}

func NewService(baseURL string, store storage.Store) *Service {
	return &Service{baseURL: strings.TrimRight(baseURL, "/"), store: store}
}

func (s *Service) Match(code string) bool {
	url, err := s.store.Get(code)
	if url == "" || err != nil {
		return false
	}
	return true
}

// Shorten creates a short code for the given URL and stores the mapping
func (s *Service) Shorten(_ *http.Request, url, preferredCode string) (string, error) {
	if url == "" {
		return "", errors.New("empty url")
	}
	code := generateCode(6)
	if preferredCode != "" {
		if !validatePreferredCode(preferredCode) {
			return "", errors.New("invalid character(s) in preferred code")
		}
		code = preferredCode
	}
	for s.Match(code) {
		if preferredCode != "" {
			return "", errors.New("preferred code is already taken")
		}
		log.Printf("WARNING collision for code %s\n", code)
		code = generateCode(6)
	}
	if _, err := s.store.Save(code, url); err != nil {
		return "", err
	}
	return code, nil
}

// Resolve returns the original URL for the given code
func (s *Service) Resolve(_ *http.Request, code string) (string, error) {
	url, err := s.store.Get(code)
	if err != nil {
		return "", errors.New("not found")
	}
	return url, nil
}

func (s *Service) Stats(_ *http.Request, code string) (string, error) {
	url, err := s.store.Get(code)
	if err != nil {
		return "", errors.New("not found")
	}
	stats := linkStats{
		Code:   code,
		Target: url,
	}
	rep, err := json.Marshal(stats)
	if err != nil {
		return "", errors.New("not found")
	}
	return string(rep), nil
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

func validatePreferredCode(code string) bool {
	return !strings.ContainsFunc(code, func(r rune) bool {
		return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9')
	})
}
