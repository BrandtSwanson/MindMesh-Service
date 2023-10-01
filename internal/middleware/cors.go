// cors.go

package middleware

import (
	"github.com/rs/cors"
)

// NewCORS returns a CORS middleware with your desired configuration.
func NewCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Replace with your React app's URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}
