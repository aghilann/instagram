package routes

import (
	"instagram/internal/handlers"
	"net/http"
)

func PostRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /post/{id}", handlers.GetPostByID)
	mux.HandleFunc("GET /post/user/{user_id}", handlers.HandleGetPostsForUser)
	mux.HandleFunc("DELETE /post/{id}", handlers.DeletePost)
	mux.HandleFunc("POST /post/", handlers.HandlePostPost)

	return mux
}
