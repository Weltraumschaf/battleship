package handler

import (
	"net/http"
	"weltraumschaf.de/battleship/internal/server/service"
)

type UserHandler struct {
	users service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		users: *service.NewUserService(),
	}
}

func (h *UserHandler) AllUsers(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}
