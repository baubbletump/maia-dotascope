package rank

import (
	"time"

	"github.com/google/uuid"
)

type PlayerRankHistory struct {
	ID          uuid.UUID `json:"id"`
	AccountID   int       `json:"account_id"`
	MMR         int       `json:"mmr"`
	Recorded_at time.Time `json:"recorded_at"`
}
