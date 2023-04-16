package auth

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

func NewPayload(userId uuid.UUID, email string, duration time.Duration) (*Payload, error) {
	currentTime := time.Now()
	payload := &Payload{
		ID:        userId,
		Email:     email,
		IssuedAt:  currentTime,
		ExpiredAt: currentTime.Add(time.Hour * duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
