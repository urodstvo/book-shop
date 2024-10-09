package payments

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

type requestBody struct {
	Cardholder *string    `json:"cardholder"`
	Number     *string    `json:"card_number"`
	ExpiredAt  *time.Time `json:"expired_at"`
	Type       *string    `json:"type"`
}

func (h *Payments) AddPayment(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	tx, err := h.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var body requestBody
	err = d.Decode(&body)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Cardholder == nil {
		h.Logger.Error("holder")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Number == nil {
		h.Logger.Error("numver")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.ExpiredAt == nil {
		h.Logger.Error("expired")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Type == nil {
		h.Logger.Error("type")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createPaymentQuery := squirrel.Insert(models.Payment{}.TableName()).Columns("user_id", "card_number", "card_type", "cardholder_name", "card_expired_at").
		Values(user.Id, body.Number, body.Type, body.Cardholder, body.ExpiredAt)

	_, err = createPaymentQuery.RunWith(tx).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)

}
