package handlers

import (
	"database/sql"
	"encoding/json"
	"instagram/internal/middleware"
	"instagram/internal/models"
	"instagram/internal/repositories"
	"instagram/internal/utils"
	"net/http"
	"time"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Get the database connection from the context
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	// Decode the user from the request body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure email and password are provided
	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}

	// Hash the plaintext password
	user.PasswordHash, _ = utils.HashPassword(user.Password)

	// Get the user metadata like ID from the database if it matches the email and passwordHash
	auth, err := repositories.GetUserAuth(db, user.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare the password hash from the database with the hashed password
	isCorrectPassword := utils.VerifyPassword(user.Password, auth.PasswordHash)
	if !isCorrectPassword {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate the JWT token with the user's ID
	token, claims, err := utils.GenerateJWT(auth.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return the token and expiration time to the client
	response := map[string]interface{}{
		"token":      token,
		"expires_at": claims.ExpiresAt.Time.Format(time.RFC3339),
		"id":         auth.ID,
	}

	// Set the response headers and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate that necessary fields are provided
	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Username, Email, and Password are required", http.StatusBadRequest)
		return
	}

	// Hash the user's password before storing it in the database
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.PasswordHash = string(hashedPassword) // Store hashed password
	user.Password = ""                         // Don't store plain-text password

	// Save the user to the database
	newUser, err := repositories.SaveUser(db, &user)
	if err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a JWT token for the new user
	token, claims, err := utils.GenerateJWT(newUser.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send back the JWT token and expiration to the client
	response := map[string]interface{}{
		"token":      token,
		"expires_at": claims.ExpiresAt.Time.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
