package models

import "time"

type Follow struct {
	FollowerID  int       `json:"follower_id" db:"follower_id"`
	FollowingID int       `json:"following_id" db:"following_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
