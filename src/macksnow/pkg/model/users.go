package model

import (
	"time"
)

type User struct {
	ID                int       `db:"id"`
	Nickname          string    `validate:"required,isUTF8,min=4,max=15" db:"nickname"`
	ScreenName        string    `validate:"required,alphanum,min=4,max=15" db:"screen_name"`
	Email             string    `validate:"required,email" db:"email"`
	PasswordHash      string    `validate:"required" db:"password_hash"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	DeletedAt         time.Time `db:"deleted_at"`
	ConfirmationToken string    `db:"confirmation_token"`
	ConfirmedAt       time.Time `db:"confirmed_at"`
}
