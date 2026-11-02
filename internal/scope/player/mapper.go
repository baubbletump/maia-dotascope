package player

import "github.com/baubbletump/maia-dotascope/internal/scope/dota/models"

func PlayerToPlayer(player *models.PlayerRankResponse) Player {
	return Player{
		AccountID:  player.Profile.AccountID,
		Username:   player.Profile.PersonaName,
		ProfileUrl: player.Profile.ProfileURL,
	}
}
