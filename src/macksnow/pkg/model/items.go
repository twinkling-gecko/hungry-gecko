package model

import (
	"time"
)

type Item struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" validate:"required" db:"name"`
	Summary   string    `json:"summary" validate:"required" db:"summary"`
	URI       string    `json:"uri" validate:"required" db:"uri"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
