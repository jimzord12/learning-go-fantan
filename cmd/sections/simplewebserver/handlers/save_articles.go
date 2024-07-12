package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jimzord12/learning-go-fantan/cmd/sections/simplewebserver/fakedb"
)

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-rapid-key")

	// This usually is a middleware
	if !ValidateAPIKey(apiKey) {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Client (JSON) -> Web Server (Go)
	var art fakedb.Article
	err := json.NewDecoder(r.Body).Decode(&art)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Web Server (Go) -> DB (Go))
	err = fakedb.SaveArticle(art)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, this is a plain text response!")
}
