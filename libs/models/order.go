package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Amount    uint      `json:"amount"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}
