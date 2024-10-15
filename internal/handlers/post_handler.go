package handlers

import (
	"database/sql"
	"encoding/json"
	"instagram/internal/middleware"
	"instagram/internal/models"
	"instagram/internal/repositories"
	"net/http"
	"strconv"
)

func HandlePostPost(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.AddPost(db, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repositories.DeletePost(db, postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	post, err := repositories.GetPostByID(db, postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleGetPostsForUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	posts, err := repositories.GetPostsForUser(db, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
