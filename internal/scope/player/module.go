package player

import (
	"log/slog"
	"net/http"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(mux *http.ServeMux, pool *pgxpool.Pool, logger *slog.Logger) {
	repo := NewRepository(pool)

	dotaClient := dota.NewClient()
	service := NewService(repo, dotaClient)
	handler := NewHandler(service, logger)

	handler.RegisterRoutes(mux)
}
