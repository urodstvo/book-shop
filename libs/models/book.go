package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Cover      string    `json:"cover"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`

	Rating      uint   `json:"rating"`
	RatingCount uint   `json:"rating_count"`
	OrdersCount uint   `json:"orders_count"`
	StockCount  uint   `json:"stock_count"`
	Annotation  string `json:"annotation"`
	PageCount   uint   `json:"page_count"`

	Author      string    `json:"author"`
	PublishedBy string    `json:"published_by"`
	PublishedAt time.Time `json:"published_at"`
}
