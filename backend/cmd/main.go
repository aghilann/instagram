package main

import (
	"database/sql"
	"fmt"
	"instagram/internal/middleware"
	"instagram/internal/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "instagram.db")
	if err != nil {
		panic(err)
	}

	// Enable foreign key support in SQLite
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(err)
	}

	// Wrap the mux with the DB middleware, and then with the CORS middleware
	var muxWithMiddleware http.Handler
	muxWithMiddleware = middleware.DBMiddleware(mux, db)
	muxWithMiddleware = middleware.CORSMiddleware(muxWithMiddleware)
	muxWithMiddleware = middleware.LoggingMiddleware(muxWithMiddleware)

	// Protect /users/ and /follow/ routes with JWTMiddleware
	mux.Handle("/users/", middleware.JWTMiddleware(routes.UserRouter()))
	mux.Handle("/follow/", middleware.JWTMiddleware(routes.FollowRouter()))
	mux.Handle("/post/", middleware.JWTMiddleware(routes.PostRouter()))
	mux.Handle("/comment/", middleware.JWTMiddleware(routes.CommentRouter()))

	// Do not protect /auth/ route (for login, registration, etc.)
	mux.Handle("/auth/", routes.AuthRouter())

	fmt.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", muxWithMiddleware)
	if err != nil {
		panic(err)
	}
}
