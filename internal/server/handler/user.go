package handler

import (
	"encoding/json"
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
	err := json.NewEncoder(w).Encode(h.users.AllUsers())
	if err != nil {
		http.Error(w, "Can't encode data to JSON!", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}

func (h *UserHandler) SingleUser(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}
