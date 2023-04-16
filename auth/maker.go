package auth

import (
	"github.com/google/uuid"
	"time"
)

type Maker interface {
	CreateToken(userId uuid.UUID, username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
