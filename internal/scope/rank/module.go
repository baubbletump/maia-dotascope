package rank

import (
	"log/slog"
	"net/http"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota"
	"github.com/baubbletump/maia-dotascope/internal/scope/player"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(mux *http.ServeMux, pool *pgxpool.Pool, logger *slog.Logger) {
	repo := NewRepository(pool)
	playerRepo := player.NewRepository(pool)
	dotaClient := dota.NewClient()

	playerService := player.NewService(playerRepo, dotaClient)

	service := NewService(repo, dotaClient, playerService)
	handler := NewHandler(service, logger)

	handler.RegisterRoutes(mux)
}
