package models

import "time"

type Post struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	ImageURL  string    `json:"image_url" db:"image_url"`
	Caption   string    `json:"caption,omitempty" db:"caption"`
	CreatedAt time.Time `json:"post_created_at" db:"created_at"` //
}

type FeedPost struct {
	Post
	User
}
