package endpoints

import (
	"net/http"
	"quakelog/application/domain/quakelog"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	GameService quakelog.GameService
}

func (h *Handler) InitializeRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/v1/quake", h.GameFindAll)
	r.Get("/v1/quake/{gameId}", h.GameFindById)

	return r
}
