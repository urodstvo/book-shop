package orders

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Orders) GetDynamic(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	if start == "" || end == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getOrdersByPeriodQuery := squirrel.Select("*").From(models.Order{}.TableName()).
		Where("created_at >= ? AND created_at <= ?", start, end)

	orders := []models.Order{}
	rows, err := getOrdersByPeriodQuery.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		order := models.Order{}
		var status string
		err = rows.Scan(&order.Id, &order.UserId, &order.PaymentId, &order.Price, &status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		order.Status = status
		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
