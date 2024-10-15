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

func TestHandlePostPost(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	// Create a POST request to create a post
	post := models.Post{
		UserID:    1, // We are not enforcing foreign keys
		ImageURL:  "https://example.com/image.jpg",
		Caption:   "Test Caption",
		CreatedAt: time.Now(),
	}
	postJSON, _ := json.Marshal(post)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(postJSON))

	// Add context with DB
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandlePostPost)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify the post record is in the database
	row := db.QueryRow("SELECT user_id, image_url, caption FROM posts WHERE user_id = ? AND image_url = ?", post.UserID, post.ImageURL)
	var userID int
	var imageURL, caption string
	err := row.Scan(&userID, &imageURL, &caption)
	if err != nil {
		t.Fatalf("failed to query post: %v", err)
	}
	assert.Equal(t, 1, userID)
	assert.Equal(t, "https://example.com/image.jpg", imageURL)
	assert.Equal(t, "Test Caption", caption)
}

func TestDeletePost(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	// Insert a post into the database
	_, err := db.Exec("INSERT INTO posts (user_id, image_url, caption) VALUES (?, ?, ?)", 1, "https://example.com/delete.jpg", "Post to be deleted")
	if err != nil {
		t.Fatalf("failed to insert post: %v", err)
	}

	// Create a DELETE request to delete the post
	req := httptest.NewRequest("DELETE", "/posts/1", nil)

	// Add context with DB and set path value for post ID
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)
	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeletePost)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify the post record is deleted from the database
	row := db.QueryRow("SELECT id FROM posts WHERE id = ?", 1)
	var id int
	err = row.Scan(&id)
	assert.NotNil(t, err) // Expecting no post found
}

func TestGetPostByID(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	// Insert a post into the database
	_, err := db.Exec("INSERT INTO posts (user_id, image_url, caption) VALUES (?, ?, ?)", 1, "https://example.com/existing.jpg", "Existing Post")
	if err != nil {
		t.Fatalf("failed to insert post: %v", err)
	}

	// Create a GET request to retrieve the post by ID
	req := httptest.NewRequest("GET", "/posts/1", nil)

	// Add context with DB and set path value for post ID
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)
	req.SetPathValue("id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPostByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var post models.Post
	err = json.NewDecoder(rr.Body).Decode(&post)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	assert.Equal(t, 1, post.UserID)
	assert.Equal(t, "https://example.com/existing.jpg", post.ImageURL)
	assert.Equal(t, "Existing Post", post.Caption)
}

func TestHandleGetPostsForUser(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	// Insert posts for a user into the database
	_, err := db.Exec("INSERT INTO posts (user_id, image_url, caption) VALUES (?, ?, ?)", 1, "https://example.com/first.jpg", "User's First Post")
	_, err = db.Exec("INSERT INTO posts (user_id, image_url, caption) VALUES (?, ?, ?)", 1, "https://example.com/second.jpg", "User's Second Post")
	if err != nil {
		t.Fatalf("failed to insert posts: %v", err)
	}

	// Create a GET request to retrieve posts for a user
	req := httptest.NewRequest("GET", "/users/1/posts", nil)

	// Add context with DB and set path value for user ID
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)
	req.SetPathValue("user_id", "1")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleGetPostsForUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var posts []models.Post
	err = json.NewDecoder(rr.Body).Decode(&posts)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	assert.Len(t, posts, 2)
	assert.Equal(t, "User's First Post", posts[0].Caption)
	assert.Equal(t, "User's Second Post", posts[1].Caption)
	assert.Equal(t, "https://example.com/first.jpg", posts[0].ImageURL)
	assert.Equal(t, "https://example.com/second.jpg", posts[1].ImageURL)
}
