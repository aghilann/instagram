package routes

import (
	"instagram/internal/handlers"
	"net/http"
)

func AuthRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/signup", handlers.HandleSignup)
	mux.HandleFunc("POST /auth/login", handlers.HandleLogin)
	return mux
}
