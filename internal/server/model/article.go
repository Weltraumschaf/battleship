package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Article struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
}

func (a Article) String() string {
	return fmt.Sprintf(
		"{Id:%s, Title:%s, Description:%s, Content:%s}",
		a.Id, a.Title, a.Description, a.Content)
}
