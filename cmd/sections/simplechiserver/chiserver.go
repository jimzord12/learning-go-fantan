package simplechiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jimzord12/learning-go-fantan/cmd/sections/simplechiserver/mymiddleware"
	"github.com/jimzord12/learning-go-fantan/cmd/sections/simplewebserver/handlers"
)

func Main() {
	const PORT string = "8080"

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Custom Middleware
	r.Use(mymiddleware.ApiKeyValidator)

	r.Get("/api/v1/articles", handlers.GetArticles)
	r.Post("/api/v1/articles", handlers.SaveArticle)

	fmt.Println("Server listens on: ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
