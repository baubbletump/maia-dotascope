package rank

import "github.com/baubbletump/maia-dotascope/internal/scope/dota/models"

func PlayerToHistory(player *models.PlayerRankResponse) PlayerRankHistory {
	return PlayerRankHistory{
		AccountID: player.Profile.AccountID,
		Tier:      player.RankTier,
	}
}
