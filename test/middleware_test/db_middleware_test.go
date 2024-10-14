package middleware_test

import (
	"database/sql"
	"instagram/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// TestDBMiddleware tests that the database connection is correctly injected into the request context
func TestDBMiddleware(t *testing.T) {
	// Create an in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to create in-memory DB: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Fatalf("failed to close DB: %v", err)
		}
	}(db)

	// Create a simple table for testing
	_, err = db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	// Create a mock handler that checks if the DB is injected
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the DB from the context
		dbFromContext, ok := middleware.GetDBFromContext(r.Context())
		if !ok {
			http.Error(w, "DB not found in context", http.StatusInternalServerError)
			return
		}

		// Ensure the DB is not nil
		if dbFromContext == nil {
			http.Error(w, "DB is nil", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("DB injected successfully"))
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	})

	// Wrap the mock handler with the DBMiddleware
	handlerWithMiddleware := middleware.DBMiddleware(mockHandler, db)

	// Create a test HTTP request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	// Serve the request
	handlerWithMiddleware.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", w.Code)
	}

	// Check the response body
	expected := "DB injected successfully"
	if w.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, w.Body.String())
	}
}
