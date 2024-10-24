package carts

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Carts) GetItems(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	items := []struct {
		models.Cart `json:"item"`
		BookName    string  `json:"book_name"`
		BookAuthor  string  `json:"book_author"`
		BookPrice   float32 `json:"book_price"`
		StockCount  uint    `json:"stock_count"`
	}{}

	getItemsQuery := squirrel.Select("c.*, b.stock_count, b.name, b.author, b.price").From(models.Cart{}.TableName() + " c").Join(models.Book{}.TableName() + " b ON c.book_id = b.id").Where(squirrel.Eq{"c.user_id": user.Id})

	rows, err := getItemsQuery.RunWith(h.DB).Query()
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		item := models.Cart{}
		var stock_count uint
		var book_name string
		var book_author string
		var book_price float32
		err = rows.Scan(&item.UserId, &item.BookId, &item.Quantity, &item.Price, &item.CreatedAt, &item.UpdatedAt, &stock_count, &book_name, &book_author, &book_price)
		if err != nil {
			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		items = append(items, struct {
			models.Cart `json:"item"`
			BookName    string  `json:"book_name"`
			BookAuthor  string  `json:"book_author"`
			BookPrice   float32 `json:"book_price"`
			StockCount  uint    `json:"stock_count"`
		}{
			Cart:       item,
			StockCount: stock_count,
			BookName:   book_name,
			BookAuthor: book_author,
			BookPrice:  book_price,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}
