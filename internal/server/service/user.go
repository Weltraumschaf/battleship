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

func (s *UserService) AllUsers() []model.User {
	return s.data.FindAll()
}

func (s *UserService) SingleUser(name string) (model.User, bool) {
	return s.data.FindByName(name)
}
