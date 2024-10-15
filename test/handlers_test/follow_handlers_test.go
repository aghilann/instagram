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

func TestHandlePostFollow(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close db: %v", err)
		}
	}(db)

	// Create a POST request to follow a user
	follow := models.Follow{
		FollowerID:  1, // Assuming user1 has ID 1
		FollowingID: 2, // Assuming user2 has ID 2
		CreatedAt:   time.Now(),
	}
	followJSON, _ := json.Marshal(follow)
	req := httptest.NewRequest("POST", "/follow", bytes.NewBuffer(followJSON))

	// Add context with DB
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandlePostFollow)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify the follow record is in the database
	row := db.QueryRow("SELECT follower_id, following_id FROM follows WHERE follower_id = ? AND following_id = ?", follow.FollowerID, follow.FollowingID)
	var followerID, followingID int
	err := row.Scan(&followerID, &followingID)
	if err != nil {
		t.Fatalf("failed to query follow: %v", err)
	}
	assert.Equal(t, 1, followerID)
	assert.Equal(t, 2, followingID)
}

func TestHandleDeleteFollow(t *testing.T) {
	db := setupTestDB(t)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close db: %v", err)
		}
	}(db)

	// Insert a follow record into the database
	_, err := db.Exec("INSERT INTO follows (follower_id, following_id) VALUES (?, ?)", 1, 2)
	if err != nil {
		t.Fatalf("failed to insert follow: %v", err)
	}

	// Create a DELETE request to unfollow a user
	follow := models.Follow{
		FollowerID:  1,
		FollowingID: 2,
	}
	followJSON, _ := json.Marshal(follow)
	req := httptest.NewRequest("DELETE", "/unfollow", bytes.NewBuffer(followJSON))

	// Add context with DB
	ctx := context.WithValue(req.Context(), middleware.DBContextKey, db)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleDeleteFollow)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify the follow record is deleted from the database
	row := db.QueryRow("SELECT follower_id, following_id FROM follows WHERE follower_id = ? AND following_id = ?", follow.FollowerID, follow.FollowingID)
	err = row.Scan(&follow.FollowerID, &follow.FollowingID)
	assert.NotNil(t, err)
}
