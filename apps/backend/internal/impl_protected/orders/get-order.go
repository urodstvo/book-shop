package orders

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Orders) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["OrderId"]

	query := squirrel.Select("*").From(models.Order{}.TableName()).Where(squirrel.Eq{"id": orderId})

	order := models.Order{}
	err := query.RunWith(h.DB).QueryRow().Scan(&order.Id, &order.UserId, &order.PaymentId, &order.Price, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payment := models.Payment{}
	getPaymentQuery := squirrel.Select("*").From(models.Payment{}.TableName()).Where(squirrel.Eq{"id": order.PaymentId})
	err = getPaymentQuery.RunWith(h.DB).QueryRow().Scan(&payment.Id, &payment.UserId, &payment.CardNumber, &payment.CardType,
		&payment.CardholderName, &payment.CardExpiredAt, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	books := []models.Book{}
	getOrderedBooksQuery := squirrel.Select("book_id").From(models.OrderBook{}.TableName()).Where(squirrel.Eq{"order_id": order.Id})
	bookRows, err := getOrderedBooksQuery.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for bookRows.Next() {
		orderedBook := models.OrderBook{}
		err = bookRows.Scan(&orderedBook.BookId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var tmp_rating sql.NullFloat64
		book := models.Book{Rating: new(float32)}
		getBookQuery := squirrel.Select("*").From(models.Book{}.TableName()).Where(squirrel.Eq{"id": orderedBook.BookId})

		err = getBookQuery.RunWith(h.DB).QueryRow().Scan(&book.Id, &book.Name, &book.Cover, &book.Author,
			&tmp_rating, &book.RatingCount, &book.Annotation, &book.Price, &book.PageCount, &book.StockCount,
			&book.OrdersCount, &book.PublishedBy, &book.PublishedAt, &book.Created_at, &book.Updated_at)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if tmp_rating.Valid {
			*book.Rating = float32(tmp_rating.Float64)
		} else {
			book.Rating = nil
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(OrderResponse{Order: &order, Books: books, Payment: payment})
}