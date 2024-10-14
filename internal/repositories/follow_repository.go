package repositories

import (
	"database/sql"
	"fmt"
	"instagram/internal/models"
)

func AddFollow(db *sql.DB, follow *models.Follow) error {
	query := `INSERT INTO follows (follower_id, following_id) VALUES (?, ?)`
	_, err := db.Exec(query, follow.FollowerID, follow.FollowingID)
	if err != nil {
		return fmt.Errorf("failed to add follow: %w", err)
	}
	return nil
}

func RemoveFollow(db *sql.DB, follow *models.Follow) error {
	query := `DELETE FROM follows WHERE follower_id = ? AND following_id = ?`

	// Execute the delete query and check the number of affected rows
	result, err := db.Exec(query, follow.FollowerID, follow.FollowingID)
	if err != nil {
		return fmt.Errorf("failed to remove follow: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no follow relationship found between follower %d and following %d", follow.FollowerID, follow.FollowingID)
	}

	return nil
}
