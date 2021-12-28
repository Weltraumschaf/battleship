package service

import (
	"weltraumschaf.de/battleship/internal/server/model"
	"weltraumschaf.de/battleship/internal/server/repo"
)

type UserService struct {
	data *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s UserService) AllUsers() []model.User {
	return s.data.FindAll()
}
