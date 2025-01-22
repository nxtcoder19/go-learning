package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Enhanced Go HTTP Server!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 - Page not found!")
}

func main() {
	mux := http.NewServeMux() // Create a new router

	// Register handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)

	// Custom 404 handler
	mux.HandleFunc("/404", notFoundHandler)

	// Start the server with logging
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting enhanced server on port 8080...")
	log.Fatal(server.ListenAndServe()) // Start the server and log errors
}
