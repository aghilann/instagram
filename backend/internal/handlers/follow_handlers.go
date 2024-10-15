package handlers

import (
	"database/sql"
	"encoding/json"
	"instagram/internal/middleware"
	"instagram/internal/models"
	"instagram/internal/repositories"
	"net/http"
)

func HandlePostFollow(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	var follow models.Follow
	err := json.NewDecoder(r.Body).Decode(&follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if follow.FollowerID == follow.FollowingID {
		http.Error(w, "You can't follow yourself", http.StatusBadRequest)
		return
	}

	if follow.FollowerID == 0 || follow.FollowingID == 0 {
		http.Error(w, "Invalid follower or following ID", http.StatusBadRequest)
		return
	}

	err = repositories.AddFollow(db, &follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleDeleteFollow(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	var follow models.Follow
	err := json.NewDecoder(r.Body).Decode(&follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if follow.FollowerID == 0 || follow.FollowingID == 0 {
		http.Error(w, "Invalid follower or following ID", http.StatusBadRequest)
		return
	}

	err = repositories.RemoveFollow(db, &follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
