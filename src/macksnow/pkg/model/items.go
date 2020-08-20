package model

import (
	"time"
)

type Item struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Summary   string    `json:"summary" validate:"required"`
	URI       string    `json:"uri" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
