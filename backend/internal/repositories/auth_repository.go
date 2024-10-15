package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"instagram/internal/models"
)

func GetUserAuth(db *sql.DB, email string) (*models.Auth, error) {
	var auth models.Auth

	query := `
        SELECT id, username, email, password_hash
        FROM users
        WHERE email = ?
    `

	err := db.QueryRow(query, email).Scan(
		&auth.ID,
		&auth.Username,
		&auth.Email,
		&auth.PasswordHash,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.Auth{}, fmt.Errorf("user with email %s not found", email)
		}
		return &models.Auth{}, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return &auth, nil
}
