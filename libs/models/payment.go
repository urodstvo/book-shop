package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id            uuid.UUID
	UserId        uuid.UUID
	CardNumber    string
	CardType      string
	CardName      string
	CardExpiredAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
