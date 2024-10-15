package handlers_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"instagram/internal/handlers"
	"instagram/internal/middleware"
	"instagram/internal/models"
	"instagram/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandleSignup(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	user := models.User{
		Auth: models.Auth{
			ID:       1,
			Username: "testuser",
			Email:    "testuser@example.com",
			Password: "password123",
		},
		Bio:          "",
		ProfileImage: "",
		CreatedAt:    time.Time{},
	}
	userJSON, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(userJSON))

	// Add context with DB
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleSignup)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify the user was inserted into the database
	row := db.QueryRow("SELECT username, email FROM users WHERE email = ?", user.Email)
	var username, email string
	err := row.Scan(&username, &email)
	if err != nil {
		t.Fatalf("failed to query inserted user: %v", err)
	}
	assert.Equal(t, "testuser", username)
	assert.Equal(t, "testuser@example.com", email)

	// Check if the response contains a valid token
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Contains(t, response, "token")
	assert.Contains(t, response, "expires_at")
}

func TestHandleLogin(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	// Insert a user into the database
	hashedPassword, _ := utils.HashPassword("password123")
	_, err := db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", "testuser", "testuser@example.com", hashedPassword)
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	// Create a login request
	loginRequest := models.User{
		Auth: models.Auth{
			ID:           0,
			Username:     "testuser",
			Email:        "testuser@example.com",
			PasswordHash: "",
			Password:     "password123",
		},
	}
	loginJSON, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))

	// Add context with DB
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleLogin)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Check if the response contains a valid token
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Contains(t, response, "token")
	assert.Contains(t, response, "expires_at")
}

func TestHandleLoginInvalidCredentials(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	// Insert a user into the database
	hashedPassword, _ := utils.HashPassword("password123")
	_, err := db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", "testuser", "testuser@example.com", hashedPassword)
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	// Create a login request with an incorrect password
	loginRequest := models.User{
		Auth: models.Auth{
			ID:           0,
			Username:     "",
			Email:        "testuser@example.com",
			PasswordHash: "",
			Password:     "wrongpassword",
		},
	}
	loginJSON, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))

	// Add context with DB
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleLogin)
	handler.ServeHTTP(rr, req)

	// Expect unauthorized status due to incorrect password
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
