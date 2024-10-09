package models

import (
	"time"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Cover  string `json:"cover"`
	Author string `json:"author"`

	Rating      *float32 `json:"rating"`
	RatingCount uint     `json:"rating_count"`
	Annotation  string   `json:"annotation"`

	Price       float32 `json:"price"`
	PageCount   uint    `json:"page_count"`
	OrdersCount uint    `json:"orders_count"`
	StockCount  uint    `json:"stock_count"`

	PublishedBy string    `json:"published_by"`
	PublishedAt time.Time `json:"published_at"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func (Book) TableName() string {
	return "books"
}
