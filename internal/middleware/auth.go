package middleware

import (
	"net/http"
	"os"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		check := "Bearer " + os.Getenv("API_KEY")

		if r.Header.Get("Authorization") != check {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
