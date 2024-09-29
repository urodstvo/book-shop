package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	UserId    uuid.UUID `json:"user_id"`
	BookId    uuid.UUID `json:"book_id"`
	Amount    uint      `json:"amount"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
