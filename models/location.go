package models

import (
	"github.com/google/uuid"
	"time"
)

type Location struct {
	ID                      uuid.UUID `json:"id"`
	Car                     Car       `json:"car"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	CurrentLocationDatetime time.Time `json:"current_location_datetime"`
}
