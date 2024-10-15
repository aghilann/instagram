package handlers_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"instagram/internal/handlers"
	"instagram/internal/middleware"
	"instagram/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Set up an in-memory SQLite DB for testing
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	// Initialize schema for testing
	_, err = db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		bio TEXT DEFAULT '',
		profile_image TEXT DEFAULT '',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		t.Fatalf("failed to create table users: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE follows (
		follower_id INTEGER NOT NULL REFERENCES users(id),
		following_id INTEGER NOT NULL REFERENCES users(id),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (follower_id, following_id)
	)`)
	if err != nil {
		t.Fatalf("failed to create table follows: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		image_url TEXT NOT NULL,
		caption TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)`)
	if err != nil {
		t.Fatalf("failed to create posts posts: %v", err)
	}
	return db
}

func TestHandleGetUsers(t *testing.T) {
	// This test is basic as your handler just returns a static string
	req := httptest.NewRequest("GET", "/users", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleGetUsers)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "GET /users", rr.Body.String())
}

func TestHandleGetUserById(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database: %v", err)
		}
	}(db)

	// Insert a user into the database
	auth := models.Auth{
		ID:           0,
		Username:     "tester",
		Email:        "tester@gmail.com",
		PasswordHash: "hashed password",
		Password:     "password",
	}
	_, err := db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)",
		auth.Username, auth.Email, auth.PasswordHash)
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	req := httptest.NewRequest("GET", "/users/", nil)
	req.SetPathValue("id", "1") // users/1

	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleGetUserById)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fetchedUser models.User
	err = json.NewDecoder(rr.Body).Decode(&fetchedUser)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	assert.Equal(t, "tester", fetchedUser.Username)
	assert.Equal(t, "tester@gmail.com", fetchedUser.Email)
}

func TestHandleDeleteUserById(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database: %v", err)
		}
	}(db)

	// Insert a user into the database
	_, err := db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", "testuser", "testuser@example.com", "hashedpassword")
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add context with DB and ID path value
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)
	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleDeleteUserById)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	// Verify the user is deleted
	row := db.QueryRow("SELECT id FROM users WHERE id = ?", 1)
	var id int
	err = row.Scan(&id)
	assert.NotNil(t, err) // Expecting no user found
}

func TestHandlePatchUser(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database: %v", err)
		}
	}(db)

	// Insert a user into the database
	_, err := db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", "testuser", "testuser@example.com", "hashedpassword")
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	// Create a PATCH request to update the username
	updatedUser := models.User{
		Auth: models.Auth{
			ID:           1,
			Username:     "updateduser",
			Email:        "tester@gmail.com",
			Password:     "password",
			PasswordHash: "hashed password",
		},
		Bio:          "bio",
		ProfileImage: "profilepic",
		CreatedAt:    time.Now(),
	}
	userJSON, _ := json.Marshal(updatedUser)
	req, err := http.NewRequest("PATCH", "/users/1", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Add context with DB and ID path value
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)
	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandlePatchUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var updatedUserResponse models.User
	err = json.NewDecoder(rr.Body).Decode(&updatedUserResponse)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	assert.Equal(t, "updateduser", updatedUserResponse.Username)

	// Verify the username is updated in the database
	row := db.QueryRow("SELECT username FROM users WHERE id = ?", 1)
	var username string
	err = row.Scan(&username)
	if err != nil {
		t.Fatalf("failed to query updated user: %v", err)
	}
	assert.Equal(t, "updateduser", username)
}
