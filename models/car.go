package models

import "github.com/google/uuid"

type Car struct {
	ID               uuid.UUID `json:"id"`
	Model            string    `json:"model"`
	Year             int       `json:"year"`
	LicensePlate     string    `json:"license_plate"`
	CurrentKm        float64   `json:"current_km"`
	MaxKm            float64   `json:"max_km"`
	FuelType         string    `json:"fuel_type"`
	HirePrice        float64   `json:"hire_price"`
	HireAvailability bool      `json:"hire_availability"`
}
