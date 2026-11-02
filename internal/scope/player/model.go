package player

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("error not found")
)

type Player struct {
	ID         uuid.UUID `json:"uuid"`
	AccountID  int       `json:"account_id"`
	Username   string    `json:"username"`
	ProfileUrl string    `json:"profile_url"`
	CreatedAt  time.Time `json:"created_at"`
}
