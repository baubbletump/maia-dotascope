package rank

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Create(ctx context.Context, rank PlayerRankHistory) (uuid.UUID, error) {
	var id uuid.UUID

	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO scope.player_rank_history (account_id, mmr) VALUES($1, $2) RETURNING id`,
		rank.AccountID,
		rank.MMR,
	).Scan(&id)

	return id, err
}

func (r *Repository) List(ctx context.Context, playerID int) ([]PlayerRankHistory, error) {
	rows, err := r.pool.Query(
		ctx,
		`SELECT * FROM scope.player_rank_history prh 
		 WHERE prh.account_id = $1
		 ORDER BY recorded_at DESC`,
		playerID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	ranks := make([]PlayerRankHistory, 0)

	for rows.Next() {
		var r PlayerRankHistory
		if err := rows.Scan(&r.ID, &r.AccountID, &r.MMR, &r.Recorded_at); err != nil {
			return nil, err
		}
		ranks = append(ranks, r)
	}
	return ranks, rows.Err()
}
