package carts

import (
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Carts) DeleteItem(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)
	book_id := mux.Vars(r)["BookId"]

	deleteCartItemQuery := squirrel.Delete(models.Cart{}.TableName()).Where(squirrel.Eq{"user_id": user.Id, "book_id": book_id})
	result, err := deleteCartItemQuery.RunWith(h.DB).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
