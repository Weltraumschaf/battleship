package model

import "github.com/google/uuid"

type Game struct {
	Id uuid.UUID `json:"id"`
}
