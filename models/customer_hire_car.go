package models

import (
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/google/uuid"
)

type CustomerHireCar struct {
	ID         uuid.UUID `json:"id"`
	Customer   Customer  `json:"customer"`
	Car        Car       `json:"car"`
	HireDate   date.Date `json:"hire_date"`
	ReturnDate date.Date `json:"return_date"`
}
