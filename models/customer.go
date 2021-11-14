package models

import "github.com/google/uuid"

type Customer struct {
	ID                   uuid.UUID `json:"id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Nationality          string    `json:"nationality"`
	IdentificationNumber string    `json:"identification_number"`
	IdentificationType   string    `json:"identification_type"`
}
