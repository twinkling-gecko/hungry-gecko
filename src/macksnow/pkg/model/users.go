package model

import (
	"time"
)

type User struct {
	ID                int       `json:"id" db:"id"`
	Nickname          string    `json:"nickname" validate:"required" db:"nickname"`
	ScreenName        string    `json:"screen_name" validate:"required" db:"screen_name"`
	Email             string    `json:"email" validate:"required" db:"email"`
	PasswordHash      string    `validate:"required" db:"password_hash"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at" db:"deleted_at"`
	ConfirmationToken string    `json:"confirmation_token" db:"confirmation_token"`
	ConfirmedAt       time.Time `json:"confirmed_at" db:"confirmed_at"`
}
