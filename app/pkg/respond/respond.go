package respond

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON response with status code
 func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
 }

 // Error writes an error response
 func Error(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, map[string]string{"error": msg})
 }
