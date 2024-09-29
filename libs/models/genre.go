package models

import "github.com/google/uuid"

type Genre struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
