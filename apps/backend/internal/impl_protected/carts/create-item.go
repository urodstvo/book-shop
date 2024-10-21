package carts

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

type createBody struct {
	BookId *int `json:"book_id"`
}

func (h *Carts) CreateItem(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var body createBody
	err := d.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.BookId == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getBookPriceQuery := squirrel.Select("price").From(models.Book{}.TableName()).Where(squirrel.Eq{"id": *body.BookId})

	var bookPrice float32
	err = getBookPriceQuery.RunWith(h.DB).QueryRow().Scan(&bookPrice)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createCartItemQuery := squirrel.Insert(models.Cart{}.TableName()).Columns("user_id", "book_id", "quantity", "price").
		Values(user.Id, body.BookId, 1, bookPrice)

	_, err = createCartItemQuery.RunWith(h.DB).Exec()
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
