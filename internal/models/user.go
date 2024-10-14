package models

import (
	"time"
)

type Auth struct {
	ID           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"-" db:"password_hash"` // PasswordHash is stored in DB but not exposed in JSON
}

type User struct {
	Auth
	Password     string    `json:"password,omitempty" db:"-"` // Password is optional in JSON, but not stored in the DB
	Bio          string    `json:"bio,omitempty" db:"bio"`
	ProfileImage string    `json:"profile_image,omitempty" db:"profile_image"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
