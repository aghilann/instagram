package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Pass request to the next handler
		next.ServeHTTP(w, r)

		// Log the time taken to process the request
		log.Printf("Completed in %s", time.Since(start))
	})
}
