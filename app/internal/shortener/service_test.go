package shortener

import "testing"

func TestGenerateCodeLength(t *testing.T) {
	for _, n := range []int{4, 6, 8, 12} {
		code := generateCode(n)
		if len(code) != n {
			t.Fatalf("expected length %d, got %d (%q)", n, len(code), code)
		}
	}
}
