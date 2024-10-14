package handlers

import (
	"database/sql"
	"encoding/json"
	"instagram/internal/middleware"
	"instagram/internal/models"
	"instagram/internal/repositories"
	"instagram/internal/utils"
	"net/http"
	"strconv"
)

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET /users"))
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
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

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Username, Email, and Password are required", http.StatusBadRequest)
		return
	}

	// Set the Hashed Password
	user.PasswordHash, err = utils.HashPassword(user.Password)
	user.Password = "" // Clear the password from memory so it's never accidentally exposed

	savedUser, err := repositories.SaveUser(db, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	savedUser.PasswordHash = "" // Clear the password hash from the response

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(savedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	user, err := repositories.GetUserByID(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, ok := r.Context().Value(middleware.DBContextKey).(*sql.DB)
	if !ok {
		http.Error(w, "Database not found", http.StatusInternalServerError)
		return
	}

	err = repositories.DeleteUserByID(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HandlePatchUser(w http.ResponseWriter, r *http.Request) {
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

	if user.Username == "" && user.Email == "" && user.Password == "" {
		http.Error(w, "At least one field is required", http.StatusBadRequest)
		return
	}

	user.Password = "" // Password cannot be updated via PATCH

	updatedUser, err := repositories.UpdateUser(db, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedUser.PasswordHash = "" // Clear the password hash from the response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
