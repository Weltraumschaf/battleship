package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Game struct {
	Id uuid.UUID `json:"id"`
}

func (g Game) String() string {
	return fmt.Sprintf("{ID:%s}", g.Id)
}
