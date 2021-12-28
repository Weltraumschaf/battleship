package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"weltraumschaf.de/battleship/internal/server/model"
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

func (h *UserHandler) SingleUser(w http.ResponseWriter, r *http.Request) {
	name, err := retrievePathParameter(r, "name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article, exists := h.users.SingleUser(name)

	if !exists {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, "can't encode data to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read request body", http.StatusInternalServerError)
		return
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "can't unmarshal request body to JSON", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(h.users.CreateUser(user))
	if err != nil {
		http.Error(w, "can't marshal data to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	name, err := retrievePathParameter(r, "name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(h.users.DeleteUser(name))
	if err != nil {
		http.Error(w, "can't encode data to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	panic("Not implemented!")
}

