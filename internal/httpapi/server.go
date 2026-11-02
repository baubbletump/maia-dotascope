package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/baubbletump/maia-dotascope/internal/scope/player"
	"github.com/baubbletump/maia-dotascope/internal/scope/rank"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

func NewServer(pool *pgxpool.Pool, logger *slog.Logger) *Server {
	return &Server{
		pool:   pool,
		logger: logger,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	rank.Register(mux, s.pool, s.logger)
	player.Register(mux, s.pool, s.logger)

	return mux
}
