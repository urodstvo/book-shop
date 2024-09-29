package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	CardNumber    string    `json:"card_number"`
	CardType      string    `json:"card_type"`
	CardName      string    `json:"card_name"`
	CardExpiredAt time.Time `json:"card_expired_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
