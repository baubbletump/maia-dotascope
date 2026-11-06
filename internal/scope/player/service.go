package player

import (
	"context"
	"errors"
	"fmt"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota"
	"github.com/baubbletump/maia-dotascope/internal/workerpool"
)

type Store interface {
	GetByAccountID(context.Context, int) (Player, error)
	List(context.Context) ([]Player, error)
	Create(context.Context, Player) (Player, error)
}

type Service struct {
	store  Store
	client dota.Client
}

func NewService(store Store, client dota.Client) *Service {
	return &Service{
		store:  store,
		client: client,
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

	plyResp, err := s.client.GetPlayerById(ctx, accountID)
	if err != nil {
		return Player{}, err
	}

	ply = PlayerToPlayer(plyResp)
	return s.store.Create(ctx, ply)
}

func (s *Service) GetOrCreateMany(ctx context.Context, ids []int) ([]Player, error) {
	results := workerpool.Map(ctx, ids, 5, s.GetOrCreate)

	players := make([]Player, 0, len(results))
	var errs []error
	for _, r := range results {
		if r.Err != nil {
			errs = append(errs, fmt.Errorf("player %d: %w", ids[r.Index], r.Err))
			continue
		}
		players = append(players, r.Value)
	}
	return players, errors.Join(errs...)
}

func (s *Service) Ensure(ctx context.Context, p Player) (Player, error) {
	return s.store.Create(ctx, p)
}
