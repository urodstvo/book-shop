package models

import "github.com/google/uuid"

type OrderBook struct {
	OrderId uuid.UUID `json:"order_id"`
	BookId  uuid.UUID `json:"book_id"`
	Amount  uint      `json:"amount"`
	Price   uint      `json:"price"`
}
