package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	UserId    uuid.UUID
	BookId    uuid.UUID
	Amount    uint
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
