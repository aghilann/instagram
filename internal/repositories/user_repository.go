package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"instagram/internal/models"
)

func SaveUser(db *sql.DB, user *models.User) (*models.User, error) {
	username, email, passwordHash, bio, profileImage :=
		user.Username, user.Email, user.PasswordHash, user.Bio, user.ProfileImage

	query := `
        INSERT INTO users (username, email, password_hash, bio, profile_image, created_at) 
        VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
    `

	// Use db.Exec to insert the user and capture the result
	result, err := db.Exec(query, username, email, passwordHash, bio, profileImage)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	// Get the last inserted ID (assuming id is an auto-increment field)
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert id: %w", err)
	}

	// Retrieve the newly created user
	newUser, err := GetUserByID(db, int(lastInsertID))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve new user: %w", err)
	}

	return newUser, nil
}

func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	var user models.User

	query := `
        SELECT id, username, email, password_hash, bio, profile_image, created_at
        FROM users
        WHERE id = ?
    `

	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Bio,
		&user.ProfileImage,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return &user, nil
}

func DeleteUserByID(db *sql.DB, id int) error {
	query := `
		DELETE FROM users
		WHERE id = ?
	`

	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}

func UpdateUser(db *sql.DB, user *models.User) (*models.User, error) {
	// Ensure the user ID is provided
	if user.ID == 0 {
		return nil, errors.New("user ID is required for updating")
	}

	// Prepare the update query
	query := `
        UPDATE users 
        SET username = ?, email = ?, bio = ?, profile_image = ? 
        WHERE id = ?
    `

	// Execute the update
	result, err := db.Exec(query, user.Username, user.Email, user.Bio, user.ProfileImage, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Check if any rows were affected (if no rows, the user may not exist)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("user with id %d not found", user.ID)
	}

	// Retrieve the updated user
	updatedUser, err := GetUserByID(db, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated user: %w", err)
	}

	return updatedUser, nil
}
