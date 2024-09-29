package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Amount    uint
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}
