package orders

import (
	"database/sql"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Orders) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)
	orderId := mux.Vars(r)["OrderId"]

	tx, err := h.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := squirrel.Select("*").From(models.Order{}.TableName()).Where(squirrel.Eq{"id": orderId})
	var status string
	order := models.Order{}
	err = query.RunWith(tx).QueryRow().Scan(&order.Id, &order.UserId, &order.PaymentId, &order.Price, &status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		tx.Rollback()
		return
	}

	order.Status = status

	if order.UserId != user.Id && user.Role != "admin" {
		w.WriteHeader(http.StatusForbidden)
		tx.Rollback()
		return
	}

	if order.Status == "delievered" || order.Status == "cancelled" {
		w.WriteHeader(http.StatusBadRequest)
		tx.Rollback()
		return
	}

	getOrderedBooksQuery := squirrel.Select("ob.*", "b.*").From(models.OrderBook{}.TableName() + " ob").Join(models.Book{}.TableName() + " b on b.id = ob.order_id").Where(squirrel.Eq{"ob.order_id": order.Id})
	bookRows, err := getOrderedBooksQuery.RunWith(tx).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	for bookRows.Next() {
		var tmp_rating sql.NullFloat64
		orderedBook := models.OrderBook{}
		book := models.Book{Rating: new(float32)}
		err = bookRows.Scan(&orderedBook.BookId, &orderedBook.OrderId, &orderedBook.Amount, &orderedBook.Price,
			&book.Id, &book.Name, &book.Cover, &book.Author,
			&tmp_rating, &book.RatingCount, &book.Annotation, &book.Price, &book.PageCount, &book.StockCount,
			&book.OrdersCount, &book.PublishedBy, &book.PublishedAt, &book.Created_at, &book.Updated_at)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		if tmp_rating.Valid {
			*book.Rating = float32(tmp_rating.Float64)
		} else {
			book.Rating = nil
		}

		book.OrdersCount--
		book.StockCount += orderedBook.Amount

		updateBookQuery := squirrel.Update(models.Book{}.TableName()).Set("orders_count", book.OrdersCount).Set("stock_count", book.StockCount).Where(squirrel.Eq{"id": book.Id})

		_, err = updateBookQuery.RunWith(tx).Exec()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tx.Rollback()
			return
		}
	}

	deleteQuery := squirrel.Delete(models.Order{}.TableName()).Where(squirrel.Eq{"id": orderId})

	_, err = deleteQuery.RunWith(tx).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}
