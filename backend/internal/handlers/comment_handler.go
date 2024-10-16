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

func HandlePostComment(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if comment.UserID == 0 || comment.PostID == 0 {
		http.Error(w, "Invalid user or post ID", http.StatusBadRequest)
		return
	}

	err = repositories.AddComment(db, &comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleGetComment(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	commentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	comments, err := repositories.GetComment(db, commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	commentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repositories.DeleteComment(db, commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleGetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	postID, err := strconv.Atoi(r.PathValue("post_id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	comments, err := repositories.GetCommentsForPost(db, postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
