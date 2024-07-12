package simplewebserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jimzord12/learning-go-fantan/cmd/sections/simplewebserver/handlers"
)

func SimpleWebServer() {
	const PORT = "8080"

	http.HandleFunc("GET /news/v2/list-by-symbol", handlers.GetArticles)
	http.HandleFunc("POST /news/v2/save-article", handlers.SaveArticle)

	fmt.Println("Server listens on:", PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	log.Fatal(err)
}
