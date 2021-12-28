package model

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (u User) String() string {
	return fmt.Sprintf("{ID:%s, Name:%s}", u.Id, u.Name)
}
