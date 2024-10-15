package routes

import (
	"instagram/internal/handlers"
	"net/http"
)

func UserRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users/", handlers.HandlePostUser)
	mux.HandleFunc("PATCH /users/", handlers.HandlePatchUser)
	mux.HandleFunc("GET /users/{id}", handlers.HandleGetUserById)
	mux.HandleFunc("DELETE /users/{id}", handlers.HandleDeleteUserById)

	return mux
}
