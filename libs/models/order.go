package models

import (
	"time"
)

type Order struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PaymentId int       `json:"payment_id"`
	Amount    uint      `json:"amount"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

func (Order) TableName() string {
	return "orders"
}
