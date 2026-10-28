// Package api serves the demo greeting.
package api

import (
	"encoding/json"
	"net/http"
)

// DefaultWebOrigin is the browser origin of the demo frontend in development.
const DefaultWebOrigin = "http://localhost:5173"

// HelloResponse is the body of GET /hello.
type HelloResponse struct {
	Message string `json:"message"`
}

// Hello answers GET /hello with the greeting.
func Hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(HelloResponse{Message: "Hello, World"})
}

// WithCORS lets the browser at origin call next. Preflight requests are
// answered here and never reach next.
func WithCORS(next http.Handler, origin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		h.Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NewServer returns the API handler.
func NewServer(webOrigin string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", Hello)
	return WithCORS(mux, webOrigin)
}
