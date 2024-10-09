package orders

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Orders) GetOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["OrderId"]

	query := squirrel.Select("status").From(models.Order{}.TableName()).Where(squirrel.Eq{"id": orderId})

	order := models.Order{}
	err := query.RunWith(h.DB).QueryRow().Scan(&order.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := struct {
		Status string `json:"status"`
	}{order.Status}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
