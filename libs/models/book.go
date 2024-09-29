package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id         uuid.UUID
	Name       string
	Created_at time.Time
	Updated_at time.Time

	Rating      uint
	OrdersCount uint
	StockCount  uint
	Annotation  string
	PageCount   uint

	Author      string
	PublishedBy string
	PublishedAt time.Time
}
