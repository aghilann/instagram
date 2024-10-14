package utils

import (
	"fmt"
	"time"
)
import "github.com/golang-jwt/jwt/v5"

var JWTSecret = []byte("an_actual_secret_instead_of_this")

// Claims structure for JWT (custom claims + standard claims)
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int) (string, *jwt.RegisteredClaims, error) {
	// Set the expiration time for the token (1 day from now)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the claims, which includes the userID and expiration time
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", nil, err
	}

	// Return the token string, registered claims, and nil for error
	return tokenString, &claims.RegisteredClaims, nil
}

// VerifyJWT Function to verify JWT tokens
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC (SigningMethodHS256 in this case)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	// Check for parsing or verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
