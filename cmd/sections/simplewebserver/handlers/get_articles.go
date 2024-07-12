package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jimzord12/learning-go-fantan/cmd/sections/simplewebserver/fakedb"
	"github.com/jimzord12/learning-go-fantan/cmd/sections/simplewebserver/restypes"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-rapid-key")

	// This usually is a middleware
	if !ValidateAPIKey(apiKey) {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	sizeStr := r.URL.Query().Get("size")
	size, _ := strconv.Atoi(sizeStr)
	if size == 0 {
		size = 10 // setting a default
	}

	// Making a "Call" to DB in order to get the Articles
	articles, err := fakedb.GetArticles(id, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the Articles to shape that the Client Expects
	var artcRes restypes.ArticleResponse
	for _, art := range articles {
		artcRes.Data = append(artcRes.Data, restypes.DataItem{
			Attributes: restypes.Attributes{
				PublishOn: art.CreatedAt,
				Title:     art.Headline,
			},
		})
	}

	// We finalize the Response, by adding a Header and
	// Converting Go -> JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artcRes)
}

func ValidateAPIKey(apiKey string) bool {
	return apiKey == "12345asd"
}
