package models

import "github.com/google/uuid"

type OrderBook struct {
	OrderId uuid.UUID
	BookId  uuid.UUID
	Amount  uint
	Price   uint
}
