package main

import (
	"encoding/json"
	"log"
	"net/http"

	"MindMesh-Service/internal/middleware" // Replace with your actual module name

	"github.com/go-chi/chi"
)

// ExampleResponse is a struct representing the JSON response.
type ExampleResponse struct {
	Sentence string `json:"sentence"`
}

func yourExampleHandler(w http.ResponseWriter, r *http.Request) {
	// Create an example response
	response := ExampleResponse{
		Sentence: "This is an example sentence from yourExampleHandler.",
	}

	// Serialize the response as JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON data to the response writer
	w.Write(jsonData)
}

func main() {
	// Create a new Chi router
	r := chi.NewRouter()

	// Apply your custom CORS middleware to the router
	r.Use(middleware.NewCORS().Handler) // Use your custom CORS middleware

	// Define a route for the example handler
	r.Get("/api/example", yourExampleHandler)

	// Start the HTTP server
	port := ":8080"
	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
