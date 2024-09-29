package models

import (
	"time"
)

type Payment struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	CardNumber    string    `json:"card_number"`
	CardType      string    `json:"card_type"`
	CardName      string    `json:"card_name"`
	CardExpiredAt time.Time `json:"card_expired_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}
