package models

type OrderBook struct {
	OrderId int     `json:"order_id"`
	BookId  int     `json:"book_id"`
	Amount  uint    `json:"amount"`
	Price   float32 `json:"price"`
}

func (OrderBook) TableName() string {
	return "orders_books"
}
