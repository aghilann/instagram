package routes

import (
	"instagram/internal/handlers"
	"net/http"
)

func CommentRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment/{id}", handlers.HandleGetComment)
	mux.HandleFunc("POST /comment/", handlers.HandlePostComment)
	mux.HandleFunc("DELETE /comment/{id}", handlers.HandleDeleteComment)
	mux.HandleFunc("GET /comment/post/{post_id}", handlers.HandleGetCommentsForPost)

	return mux
}
