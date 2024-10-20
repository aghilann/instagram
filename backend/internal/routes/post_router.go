package routes

import (
	"instagram/internal/handlers"
	"net/http"
)

func PostRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /post/{id}", handlers.HandleGetPostById)
	mux.HandleFunc("GET /post/user/{user_id}", handlers.HandleGetPostsForUser)
	mux.HandleFunc("DELETE /post/{id}", handlers.HandleDeletePost)
	mux.HandleFunc("POST /post/", handlers.HandlePostPost)
	mux.HandleFunc("GET /post/feed/{user_id}", handlers.HandleGetFeedForUser)

	return mux
}
