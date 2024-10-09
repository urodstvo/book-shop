package payments

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Payments) GetPayments(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	payments := []models.Payment{}
	getPaymentsQuery := squirrel.Select("*").From(models.Payment{}.TableName()).Where(squirrel.Eq{"user_id": user.Id})
	rows, err := getPaymentsQuery.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		payment := models.Payment{}
		err = rows.Scan(&payment.Id, &payment.UserId, &payment.CardNumber, &payment.CardType,
			&payment.CardholderName, &payment.CardExpiredAt, &payment.CreatedAt, &payment.UpdatedAt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		payments = append(payments, payment)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)

}
