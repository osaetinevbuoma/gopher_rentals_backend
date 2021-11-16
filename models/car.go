package models

import "github.com/google/uuid"

type Car struct {
	ID               uuid.UUID  `json:"id"`
	Model            string     `json:"model"`
	Year             int        `json:"year"`
	LicensePlate     string     `json:"license_plate"`
	CurrentKm        float64    `json:"current_km"`
	MaxKg            float64    `json:"max_kg"`
	FuelType         string     `json:"fuel_type"`
	HirePrice        float64    `json:"hire_price"`
	HireAvailability bool       `json:"hire_availability"`
	Locations        []Location `json:"locations"`
}
