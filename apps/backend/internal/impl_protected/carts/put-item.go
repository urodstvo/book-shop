package carts

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

type putBody struct {
	Quantity *int `json:"quantity"`
}

func (h *Carts) PutItem(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)
	book_id := mux.Vars(r)["BookId"]

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var body putBody
	err := d.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Quantity == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getStockCountQuery := squirrel.Select("stock_count").From(models.Book{}.TableName()).Where(squirrel.Eq{"id": book_id})
	var stock_count uint
	err = getStockCountQuery.RunWith(h.DB).QueryRow().Scan(&stock_count)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if *body.Quantity > int(stock_count) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateCartItemQuery := squirrel.Update(models.Cart{}.TableName()).
		Set("quantity", *body.Quantity).
		Set("updated_at", squirrel.Expr("datetime('now')")).
		Where(squirrel.Eq{"user_id": user.Id, "book_id": book_id})

	_, err = updateCartItemQuery.RunWith(h.DB).Exec()
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
