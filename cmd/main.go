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

	// Wrap the mux with the DB middleware
	muxWithDB := middleware.DBMiddleware(mux, db)

	// Protect /users/ and /follow/ routes with JWTMiddleware
	mux.Handle("/users/", middleware.JWTMiddleware(routes.UserRouter()))
	mux.Handle("/follow/", middleware.JWTMiddleware(routes.FollowRouter()))

	// Do not protect /auth/ route (for login, registration, etc.)
	mux.Handle("/auth/", routes.AuthRouter())
	// Start the server
	fmt.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", muxWithDB)
	if err != nil {
		panic(err)
	}
}
