package main

import (
	"fmt"
	"net/http"
)

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logger middleware
func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// Auth middleware
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer secret-token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// Apply middlewares to a handler
func applyMiddlewares(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func main() {
	// Sample handler that responds with "Hello, World!"
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	// Apply middlewares to the handler
	http.HandleFunc("/", applyMiddlewares(helloHandler, logger, auth))

	// Start the server on port 8080
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
