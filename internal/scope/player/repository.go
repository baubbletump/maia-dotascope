package player

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
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

func (r Repository) GetByAccountID(ctx context.Context, accountID int) (Player, error) {
	var ply Player

	err := r.pool.QueryRow(
		ctx,
		`SELECT id, account_id, username, profile_url, created_at FROM scope.player WHERE account_id = $1`,
		accountID,
	).Scan(&ply.ID, &ply.AccountID, &ply.Username, &ply.ProfileUrl, &ply.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Player{}, ErrNotFound
		}
		return Player{}, err
	}

	return ply, nil
}

// TODO: add pagination in the future
func (r Repository) List(ctx context.Context) ([]Player, error) {
	rows, err := r.pool.Query(ctx, `SELECT * FROM scope.player`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := make([]Player, 0)
	for rows.Next() {
		var ply Player
		if err := rows.Scan(&ply.ID, &ply.AccountID, &ply.Username, &ply.ProfileUrl, &ply.CreatedAt); err != nil {
			return nil, err
		}
		players = append(players, ply)
	}

	return players, nil
}

func (r Repository) Create(ctx context.Context, player Player) (Player, error) {
	var ply Player

	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO scope.player (account_id, username, profile_url)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (account_id) DO NOTHING
		 	SET username = EXCLUDED.username,
				profile_url = EXCLUDED.profile_url
		 RETURNING id, account_id, username, profile_url, created_at`,
		player.AccountID,
		player.Username,
		player.ProfileUrl,
	).Scan(
		&ply.ID,
		&ply.AccountID,
		&ply.Username,
		&ply.ProfileUrl,
		&ply.CreatedAt,
	)

	return ply, err
}
