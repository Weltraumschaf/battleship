package repo

import (
	"weltraumschaf.de/battleship/internal/server/model"
)

type UserRepo struct {
	data map[string]model.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		data: make(map[string]model.User),
	}
}

func (r *UserRepo) Save(user model.User) {
	r.data[user.Name] = user
}

func (r *UserRepo) Delete(name string) model.User {
	article, exists := r.FindByName(name)

	if exists {
		delete(r.data, name)
	}

	return article
}

func (r *UserRepo) FindAll() []model.User {
	users := make([]model.User, 0, len(r.data))

	for _, user := range r.data {
		users = append(users, user)
	}

	return users
}

func (r *UserRepo) FindByName(name string) (model.User, bool) {
	user, exists := r.data[name]
	return user, exists
}
