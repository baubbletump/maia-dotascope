package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Server struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Server {
	return &Server{logger: logger}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	return mux
}
