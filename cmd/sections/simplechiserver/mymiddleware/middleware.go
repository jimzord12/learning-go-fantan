package mymiddleware

import (
	"net/http"

	"github.com/jimzord12/learning-go-fantan/cmd/sections/simplewebserver/handlers"
)

func ApiKeyValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-rapidapi-key")

		if !handlers.ValidateAPIKey(apiKey) {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
