package repositories

import (
	"database/sql"
	"instagram/internal/models"
)

func AddComment(db *sql.DB, comment *models.Comment) error {
	_, err := db.Exec("INSERT INTO comments (user_id, post_id, content) VALUES ($1, $2, $3)", comment.UserID, comment.PostID, comment.Content)
	return err
}

func GetComment(db *sql.DB, commentID int) (*models.Comment, error) {
	var comment models.Comment
	err := db.QueryRow("SELECT id, user_id, post_id, content FROM comments WHERE id = $1", commentID).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func DeleteComment(db *sql.DB, commentID int) error {
	result, err := db.Exec("DELETE FROM comments WHERE id = $1", commentID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func GetCommentsForPost(db *sql.DB, postID int) ([]models.Comment, error) {
	rows, err := db.Query("SELECT id, user_id, post_id, content FROM comments WHERE post_id = $1", postID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
