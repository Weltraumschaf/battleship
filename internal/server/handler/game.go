package handler

import (
	"net/http"
	"weltraumschaf.de/battleship/internal/server/service"
)

type GameHandler struct {
	games service.GameService
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		games: *service.NewGameService(),
	}
}

func (h *GameHandler) AllGames(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}
