package service

import (
	"log"
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

func (s *UserService) CreateUser(user model.User) model.User {
	s.data.Save(user)
	return user
}

func (s *UserService) DeleteUser(name string) model.User {
	return s.data.Delete(name)
}

func (s *UserService) UpdateUser(toUpdate model.User) (model.User, bool) {
	log.Println("Update model: ", toUpdate)
	persisted, exists := s.data.FindByName(toUpdate.Name)

	if !exists {
		log.Println("No persistent data found!")
		return model.User{}, false
	}

	s.data.Save(persisted)
	return persisted, true
}
