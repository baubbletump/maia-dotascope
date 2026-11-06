package rank

import (
	"context"
	"time"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota"
	"github.com/baubbletump/maia-dotascope/internal/scope/player"
)

type Store interface {
	GetPlayerRankHistory(context.Context, int) ([]PlayerRankHistory, error)
	Create(context.Context, PlayerRankHistory) (PlayerRankHistory, error)
}

type Service struct {
	store         Store
	dotaClient    dota.Client
	playerService PlayerService
}

type PlayerService interface {
	Ensure(context.Context, player.Player) (player.Player, error)
}

func NewService(store Store, client dota.Client, playerService PlayerService) *Service {
	return &Service{
		store:         store,
		dotaClient:    client,
		playerService: playerService,
	}
}

func (s *Service) GetPlayerRankHistory(ctx context.Context, playerID int) ([]PlayerRankHistory, error) {
	list, err := s.store.GetPlayerRankHistory(ctx, playerID)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 && time.Since(list[0].RecordedAt) < 5*time.Minute {
		return list, nil
	}

	resp, err := s.dotaClient.GetPlayerById(ctx, playerID)
	if err != nil {
		return nil, err
	}

	rank := PlayerToHistory(resp)

	if len(list) > 0 && list[0].Tier == rank.Tier {
		return list, nil
	}

	if _, err := s.playerService.Ensure(ctx, player.PlayerToPlayer(resp)); err != nil {
		return nil, err
	}

	row, err := s.Create(ctx, rank)
	if err != nil {
		return nil, err
	}

	return append([]PlayerRankHistory{row}, list...), nil
}

func (s *Service) Create(ctx context.Context, rank PlayerRankHistory) (PlayerRankHistory, error) {
	return s.store.Create(ctx, rank)
}
