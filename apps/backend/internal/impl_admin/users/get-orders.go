package users

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

type responseItem struct {
	*models.Order
	UserName  string `json:"user_name"`
	UserLogin string `json:"user_login"`
	BookCount int    `json:"book_count"`
}

type response struct {
	Orders            []responseItem `json:"orders"`
	UniqueUsers       int            `json:"unique_users"`
	TotalPrice        float32        `json:"total_price"`
	TotalOrderedBooks int            `json:"total_ordered_books"`
}

func (h *Users) GetOrders(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	user := r.URL.Query().Get("user")

	if start == "" || end == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getOrdersByPeriodQuery := squirrel.Select("o.*, u.name, u.login").From(models.Order{}.TableName()+" o").
		Join(models.User{}.TableName()+" u ON o.user_id = u.id").Where("o.created_at >= ? AND o.created_at <= ?", start, end)

	if user != "" {
		getOrdersByPeriodQuery = getOrdersByPeriodQuery.Where(squirrel.Eq{"u.id": user})
	}

	var uniqueUsers []int = []int{}
	var totalOrderedBooks int = 0
	var totalPrice float32 = 0.0

	orders := []responseItem{}
	rows, err := getOrdersByPeriodQuery.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		order := models.Order{}
		var status string
		var userName string
		var userLogin string
		err = rows.Scan(&order.Id, &order.UserId, &order.PaymentId, &order.Price, &status, &order.CreatedAt, &order.UpdatedAt, &userName, &userLogin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		order.Status = status

		if !slices.Contains(uniqueUsers, order.UserId) {
			uniqueUsers = append(uniqueUsers, order.UserId)
		}
		totalPrice += order.Price

		getBooksCountInOrderQuery := squirrel.Select("COUNT(*)").From(models.OrderBook{}.TableName()).Where(squirrel.Eq{"order_id": order.Id})

		var bookCount int
		err = getBooksCountInOrderQuery.RunWith(h.DB).QueryRow().Scan(&bookCount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		totalOrderedBooks += bookCount

		orders = append(orders, responseItem{
			Order:     &order,
			UserName:  userName,
			UserLogin: userLogin,
			BookCount: bookCount,
		})
	}

	res := response{
		Orders:            orders,
		UniqueUsers:       len(uniqueUsers),
		TotalPrice:        totalPrice,
		TotalOrderedBooks: totalOrderedBooks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
