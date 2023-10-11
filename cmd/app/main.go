package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MindMesh-Service/internal/middleware" // Replace with your actual module name

	"github.com/go-chi/chi"
)

// ExampleResponse is a struct representing the JSON response.
type ExampleResponse struct {
	Sentence string `json:"sentence"`
}

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// Create a new Chi router
	r := chi.NewRouter()
	var notes []Note

	// Apply your custom CORS middleware to the router
	r.Use(middleware.NewCORS().Handler) // Use your custom CORS middleware

	// Create a new note
	r.Post("/api/notes", func(w http.ResponseWriter, r *http.Request) {
		var newNote Note
		err := json.NewDecoder(r.Body).Decode(&newNote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Assign a unique ID and add the note to the slice
		notes = append(notes, newNote)

		// Return the created note as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newNote)
	})

	// Retrieve all notes
	r.Get("/api/notes", func(w http.ResponseWriter, r *http.Request) {
		// Return all notes as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	})

	// Retrieve a single note by ID
	r.Get("/api/notes/{noteID}", func(w http.ResponseWriter, r *http.Request) {
		noteID := chi.URLParam(r, "noteID")
		for _, note := range notes {
			if strconv.Itoa(note.ID) == noteID {
				// Return the note as JSON
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(note)
				return
			}
		}
		// If note is not found, return a 404
		http.NotFound(w, r)
	})

	// Update a note by ID
	r.Put("/api/notes/{noteID}", func(w http.ResponseWriter, r *http.Request) {
		noteID := chi.URLParam(r, "noteID")
		var updatedNote Note
		err := json.NewDecoder(r.Body).Decode(&updatedNote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, note := range notes {
			if strconv.Itoa(note.ID) == noteID {
				// Update the note
				notes[i] = updatedNote
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		// If note is not found, return a 404
		http.NotFound(w, r)
	})

	// Delete a note by ID
	r.Delete("/api/notes/{noteID}", func(w http.ResponseWriter, r *http.Request) {
		noteID := chi.URLParam(r, "noteID")
		for i, note := range notes {
			if strconv.Itoa(note.ID) == noteID {
				// Remove the note from the slice
				notes = append(notes[:i], notes[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		// If note is not found, return a 404
		http.NotFound(w, r)
	})

	// Start the server
	http.ListenAndServe(":8080", r)

}
