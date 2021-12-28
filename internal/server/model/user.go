package model

import (
	"fmt"
)

type User struct {
	Name string    `json:"name"`
}

func (u User) String() string {
	return fmt.Sprintf("{Name:%s}", u.Name)
}
