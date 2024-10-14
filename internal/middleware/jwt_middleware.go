package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"instagram/internal/utils"
	"net/http"
	"strings"
)

// JWTMiddleware verifies the JWT token and allows the request to proceed if valid
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Split the header to get the token part
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := tokenParts[1]

		// Verify the token
		token, err := utils.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Extract the claims (e.g., userID) from the token if necessary
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Add the user ID from the claims to the request context
			userID := int(claims["user_id"].(float64))
			ctx := context.WithValue(r.Context(), "user_id", userID)
			// Proceed to the next handler with the modified context
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		}
	})
}
