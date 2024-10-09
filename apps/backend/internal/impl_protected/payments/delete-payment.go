package payments

import (
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Payments) DeletePayment(w http.ResponseWriter, r *http.Request) {
	payment_id := mux.Vars(r)["PaymentId"]
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	tx, err := h.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	deletePaymentQuery := squirrel.Delete(models.Payment{}.TableName()).Where(squirrel.Eq{"user_id": user.Id, "id": payment_id})
	sql, err := deletePaymentQuery.RunWith(tx).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	rows, err := sql.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	if rows == 0 {
		w.WriteHeader(http.StatusBadRequest)
		tx.Rollback()
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}
