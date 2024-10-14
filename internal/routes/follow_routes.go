package routes

import (
	"instagram/internal/handlers"
	"net/http"
)

func FollowRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /follow/", handlers.HandlePostFollow)
	mux.HandleFunc("DELETE /follow/", handlers.HandleDeleteFollow)

	return mux
}
