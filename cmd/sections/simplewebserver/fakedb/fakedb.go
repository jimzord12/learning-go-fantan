package fakedb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Article struct {
	Symbol    string
	CreatedAt time.Time
	Headline  string
}

func GetArticles(id string, size int) ([]Article, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	fmt.Printf("Current working directory: %s\n", cwd)

	// Get the directory of the executable
	execDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Error getting executable directory: %v", err)
	}
	fmt.Println(execDir)

	file, err := os.ReadFile(filepath.Join(execDir, "articles.json"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// From DB(Local JSON file) -> Go
	var allArticles []Article
	if err := json.Unmarshal(file, &allArticles); err != nil {
		return nil, err
	}

	// Implement Filtering
	var filteredArticles []Article

	for _, art := range allArticles {
		if len(filteredArticles) == size {
			break
		}

		if art.Symbol == id {
			filteredArticles = append(filteredArticles, art)
		}
	}

	// From Go -> JSON (expected by Client)
	return filteredArticles, nil

}

func SaveArticle(newArtc Article) error {
	// Get the directory of the executable
	execDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Error getting executable directory: %v", err)
	}
	fmt.Println(execDir)

	// Getting the JSON file that contains all the Articles
	file, err := os.ReadFile(filepath.Join(execDir, "articles.json"))
	if err != nil {
		return err
	}

	// Converting the File ([]byte) -> Go value (struct)
	// so that we can work with it
	var articles []Article
	json.Unmarshal(file, &articles)

	// We add the new Articles to the existing Articles
	articles = append(articles, newArtc)
	newFile, err := json.Marshal(articles)
	if err != nil {
		return err
	}

	os.WriteFile("newArticles.json", newFile, os.FileMode(0644))
	return nil
}
