package rank

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota"
	"github.com/baubbletump/maia-dotascope/internal/utils"
)

type Handler struct {
	service Store
	client  dota.Client
	logger  *slog.Logger
}

func NewHandler(service *Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /rank/{accountId}", h.getList)
}

func (h *Handler) getList(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(r.PathValue("accountId"), 10, 64)
	if err != nil || accountID <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "invalid account id")
		return
	}

	items, err := h.service.GetPlayerRankHistory(r.Context(), int(accountID))
	if err != nil {
		h.logger.Error("request failed", "error", err)
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if len(items) == 0 {
		utils.WriteError(w, http.StatusNotFound, "not found list by account id")
		return
	}

	utils.WriteJSON(w, http.StatusOK, items)
}
