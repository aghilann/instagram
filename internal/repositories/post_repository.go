package repositories

import (
	"database/sql"
	"fmt"
	"instagram/internal/models"
)

func AddPost(db *sql.DB, post *models.Post) error {
	query := `INSERT INTO posts (user_id, image_url, caption, created_at) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, post.UserID, post.ImageURL, post.Caption, post.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to add post: %w", err)
	}
	return nil
}

func DeletePost(db *sql.DB, postID int) error {
	query := `DELETE FROM posts WHERE id = ?`
	result, err := db.Exec(query, postID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("post with id %d not found", postID)
	}

	return nil
}

func GetPostByID(db *sql.DB, postID int) (*models.Post, error) {
	query := `SELECT id, user_id, image_url, caption, created_at FROM posts WHERE id = ?`
	row := db.QueryRow(query, postID)

	var post models.Post
	err := row.Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Caption, &post.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return &post, nil
}

func GetPostsForUser(db *sql.DB, userID int) ([]models.Post, error) {
	query := `SELECT id, user_id, image_url, caption, created_at FROM posts WHERE user_id = ?`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Caption, &post.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}
