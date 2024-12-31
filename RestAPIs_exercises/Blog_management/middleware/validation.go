package middleware

import (
	"net/http"
)

func ValidateJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Invalid Content-Type, expected application/json", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	}
}
