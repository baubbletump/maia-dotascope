package player

import (
	"context"
	"errors"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota"
)

type Store interface {
	GetByAccountID(context.Context, int) (Player, error)
	List(context.Context) ([]Player, error)
	Create(context.Context, Player) (Player, error)
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

func (s *Service) GetByAccountID(ctx context.Context, accountID int) (Player, error) {
	return s.GetOrCreate(ctx, accountID)
}

func (s *Service) List(ctx context.Context) ([]Player, error) {
	return s.store.List(ctx)
}

func (s *Service) GetOrCreate(ctx context.Context, accountID int) (Player, error) {
	ply, err := s.store.GetByAccountID(ctx, accountID)

	if !errors.Is(err, ErrNotFound) {
		return ply, err
	}

	plyResp, err := s.dotaClient.GetPlayerById(ctx, accountID)
	if err != nil {
		return Player{}, err
	}

	ply = PlayerToPlayer(plyResp)
	return s.store.Create(ctx, ply)
}

func (s *Service) Ensure(ctx context.Context, p Player) (Player, error) {
	return s.store.Create(ctx, p)
}
