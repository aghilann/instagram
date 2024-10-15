package middleware

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

const DBContextKey = "db"

// DBMiddleware injects a *sql.DB database connection into the request context.
func DBMiddleware(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure the database connection is valid
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		// Add the *sql.DB to the context
		ctx := context.WithValue(r.Context(), DBContextKey, db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetDBFromContext Helper function to retrieve the *sql.DB from the context
func GetDBFromContext(ctx context.Context) (*sql.DB, bool) {
	db, ok := ctx.Value(DBContextKey).(*sql.DB)
	return db, ok
}
