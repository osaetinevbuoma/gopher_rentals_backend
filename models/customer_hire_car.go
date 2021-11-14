package models

import (
	"github.com/google/uuid"
	"time"
)

type CustomerHireCar struct {
	ID         uuid.UUID `json:"id"`
	Customer   Customer  `json:"customer"`
	Car        Car       `json:"car"`
	HireDate   time.Time `json:"hire_date"`
	ReturnDate time.Time `json:"return_date"`
}
