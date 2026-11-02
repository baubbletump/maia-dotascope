package player

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/baubbletump/maia-dotascope/internal/utils"
)

type Handler struct {
	service *Service
	logger  *slog.Logger
}

func NewHandler(service *Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /player/{accountId}", h.getByAccountID)
	mux.HandleFunc("GET /player/", h.getList)
}

func (h *Handler) getByAccountID(w http.ResponseWriter, r *http.Request) {
	accountId, err := strconv.ParseInt(r.PathValue("accountId"), 10, 64)
	if err != nil && accountId <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "invalid player id")
		return
	}

	ply, err := h.service.GetByAccountID(r.Context(), int(accountId))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			utils.WriteError(w, http.StatusNotFound, "player not found")
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	utils.WriteJSON(w, http.StatusOK, ply)
}

func (h *Handler) getList(w http.ResponseWriter, r *http.Request) {
	players, err := h.service.List(r.Context())
	if err != nil {
		h.logger.Error("request failed", "error", err)
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	utils.WriteJSON(w, http.StatusOK, players)
}
