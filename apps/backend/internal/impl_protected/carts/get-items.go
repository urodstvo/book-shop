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
		StockCount  uint `json:"stock_count"`
	}{}

	getItemsQuery := squirrel.Select("c.*, b.stock_count").From(models.Cart{}.TableName() + " c").Join(models.Book{}.TableName() + " b ON c.book_id = b.id").Where(squirrel.Eq{"c.user_id": user.Id})

	rows, err := getItemsQuery.RunWith(h.DB).Query()
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		item := models.Cart{}
		var stock_count uint
		err = rows.Scan(&item.UserId, &item.BookId, &item.Quantity, &item.Price, &item.CreatedAt, &item.UpdatedAt, &stock_count)
		if err != nil {
			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		items = append(items, struct {
			models.Cart `json:"item"`
			StockCount  uint `json:"stock_count"`
		}{
			Cart:       item,
			StockCount: stock_count,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}
