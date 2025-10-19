package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"url-shortener/internal/config"
	"url-shortener/internal/storage/memory"
)

func TestShortenAndResolve(t *testing.T) {
	cfg := config.Load()
	store := memory.NewInMemoryStore()
	r := NewRouter(cfg, store)

	// Shorten
	body, _ := json.Marshal(map[string]string{"url": "https://golang.org"})
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("shorten status: %d, body: %s", rec.Code, rec.Body.String())
	}
	var resp struct{ Code, ShortURL string }
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Code == "" {
		t.Fatalf("expected code, got empty")
	}

	// Resolve
	req2 := httptest.NewRequest(http.MethodGet, "/"+resp.Code, nil)
	rec2 := httptest.NewRecorder()
	r.ServeHTTP(rec2, req2)
	if rec2.Code != http.StatusFound {
		t.Fatalf("resolve status: %d", rec2.Code)
	}
	loc := rec2.Header().Get("Location")
	if loc != "https://golang.org" {
		t.Fatalf("expected redirect to golang, got %s", loc)
	}
}
