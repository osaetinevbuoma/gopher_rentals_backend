package models

import (
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/google/uuid"
)

type Location struct {
	ID uuid.UUID `json:"id"`
	Car Car `json:"car"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CurrentLocationDatetime date.Date `json:"current_location_datetime"`
}
