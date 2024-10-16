package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func CORSHandler(next http.Handler) http.Handler {
	// Create a new CORS handler with permissive settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true, // Allow credentials (e.g., cookies, authorization headers)
	})

	return c.Handler(next) // Apply the CORS middleware to the next handler
}
