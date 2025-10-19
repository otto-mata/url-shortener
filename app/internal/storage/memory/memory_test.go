package memory

import "testing"

func TestInMemoryStore_SaveGet(t *testing.T) {
	s := NewInMemoryStore()
	if err := s.Save("abc", "https://example.com"); err != nil {
		t.Fatalf("save error: %v", err)
	}
	got, ok := s.Get("abc")
	if !ok || got != "https://example.com" {
		t.Fatalf("expected hit with url, got ok=%v url=%q", ok, got)
	}
	if _, ok := s.Get("nope"); ok {
		t.Fatalf("expected miss for unknown code")
	}
}
