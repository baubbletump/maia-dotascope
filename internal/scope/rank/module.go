package rank

import (
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(mux *http.ServeMux, pool *pgxpool.Pool, logger *slog.Logger) {
	repo := NewRepository(pool)
	service := NewService(repo)
	handler := NewHandler(service, logger)

	handler.RegisterRoutes(mux)
}
