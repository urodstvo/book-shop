package models

import (
	"time"
)

type Order struct {
	Id        int         `json:"id"`
	UserId    int         `json:"user_id"`
	PaymentId int         `json:"payment_id"`
	Price     float32     `json:"price"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderStatus = string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusAccepted  OrderStatus = "accepted"
	OrderStatusDelivered OrderStatus = "delivered"
)
