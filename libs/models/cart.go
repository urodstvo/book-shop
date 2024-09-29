package models

import (
	"time"
)

type Cart struct {
	UserId    int       `json:"user_id"`
	BookId    int       `json:"book_id"`
	Quantity  uint      `json:"quantity"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Cart) TableName() string {
	return "carts"
}
