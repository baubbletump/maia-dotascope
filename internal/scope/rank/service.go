package rank

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	List(context.Context, int) ([]PlayerRankHistory, error)
	Create(context.Context, PlayerRankHistory) (uuid.UUID, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) List(ctx context.Context, playerID int) ([]PlayerRankHistory, error) {
	return s.store.List(ctx, playerID)
}

func (s *Service) Create(ctx context.Context, rank PlayerRankHistory) (uuid.UUID, error) {
	return s.store.Create(ctx, rank)
}
