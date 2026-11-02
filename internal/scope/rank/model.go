package rank

import (
	"time"

	"github.com/google/uuid"
)

type PlayerRankHistory struct {
	ID         uuid.UUID `json:"id"`
	AccountID  int       `json:"account_id"`
	Tier       int       `json:"tier"`
	RecordedAt time.Time `json:"recorded_at"`
}
