package rank

import (
	"context"
	"time"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota"
)

type Store interface {
	GetPlayerRankHistory(context.Context, int) ([]PlayerRankHistory, error)
	Create(context.Context, PlayerRankHistory) (PlayerRankHistory, error)
}

type Service struct {
	store      Store
	dotaClient dota.Client
}

func NewService(store Store, client dota.Client) *Service {
	return &Service{
		store:      store,
		dotaClient: client,
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

	player, err := s.dotaClient.GetPlayerById(ctx, playerID)
	if err != nil {
		return nil, err
	}

	rank := PlayerToHistory(player)

	if len(list) > 0 && list[0].Tier == rank.Tier {
		return list, nil
	}

	// TODO: При создании новой записи без игрока, будет ругаться из-за внешнего ключа
	row, err := s.Create(ctx, rank)
	if err != nil {
		return nil, err
	}
	list = append([]PlayerRankHistory{row}, list...)

	return list, nil
}

func (s *Service) Create(ctx context.Context, rank PlayerRankHistory) (PlayerRankHistory, error) {
	return s.store.Create(ctx, rank)
}
