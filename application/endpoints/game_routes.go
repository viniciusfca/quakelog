package endpoints

import (
	"encoding/json"
	"net/http"
	"quakelog/application/contract"
	"quakelog/application/errors"

	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) GameFindAll(w http.ResponseWriter, r *http.Request) {
	games, err := h.GameService.FindAll()

	if err != nil {
		errors.SendErrorResponse(w, &errors.APIError{StatusCode: err.StatusCode, Message: err.Message})
		return
	}

	body, _ := json.Marshal(contract.ConvertGameListToGameOutList(games))
	w.Write(body)
}

func (h *Handler) GameFindById(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "gameId")
	gameId, _ := strconv.Atoi(param)
	game, err := h.GameService.FindById(gameId)

	if err != nil {
		errors.SendErrorResponse(w, &errors.APIError{StatusCode: err.StatusCode, Message: err.Message})
		return
	}

	body, _ := json.Marshal(contract.ConvertToGameOut(&game))
	w.Write(body)
}
