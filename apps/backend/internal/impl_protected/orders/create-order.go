package orders

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

type requestBody struct {
	Books []struct {
		BookId int  `json:"book_id"`
		Amount uint `json:"amount"`
	} `json:"books"`

	PaymentId int `json:"payment_id"`
}

func (h *Orders) CreateOrder(w http.ResponseWriter, r *http.Request) {
	tx, err := h.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	f := requestBody{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err = d.Decode(&f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		tx.Rollback()
		return
	}

	if len(f.Books) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookIds := []int{}
	for _, book := range f.Books {
		bookIds = append(bookIds, book.BookId)

		if book.Amount <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	getBooksQuery := squirrel.Select("id, price, stock_count, orders_count").From(models.Book{}.TableName()).Where(squirrel.Eq{"id": bookIds})
	books := []models.Book{}
	brows, err := getBooksQuery.RunWith(tx).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	totalPrice := float32(0)
	for brows.Next() {
		var book models.Book
		err = brows.Scan(&book.Id, &book.Price, &book.StockCount, &book.OrdersCount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		ind := len(books)
		if book.StockCount < f.Books[ind].Amount {
			w.WriteHeader(http.StatusBadRequest)
			tx.Rollback()
			return
		}

		totalPrice += book.Price * float32(f.Books[ind].Amount)
		books = append(books, book)
	}

	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	var orderId int
	createOrderQuery := squirrel.Insert(models.Order{}.TableName()).Columns("user_id", "payment_id", "price").Values(user.Id, f.PaymentId, totalPrice).Suffix("RETURNING id")

	err = createOrderQuery.RunWith(tx).QueryRow().Scan(&orderId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	// create order-books and update stock count
	for ind, book := range books {
		createOrdersBooksQuery := squirrel.Insert(models.OrderBook{}.TableName()).Columns("order_id", "book_id", "amount", "price").Values(orderId, book.Id, f.Books[ind].Amount, book.Price*float32(f.Books[ind].Amount))
		updateStockQuery := squirrel.Update(models.Book{}.TableName()).Set("stock_count", book.StockCount-f.Books[ind].Amount).Set("orders_count", book.OrdersCount+1).Where(squirrel.Eq{"id": book.Id})

		_, err = updateStockQuery.RunWith(tx).Exec()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tx.Rollback()
			return
		}

		_, err = createOrdersBooksQuery.RunWith(tx).Exec()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tx.Rollback()
			return
		}
	}

	clearCartQuery := squirrel.Delete(models.Cart{}.TableName()).Where(squirrel.Eq{"user_id": user.Id})
	_, err = clearCartQuery.RunWith(tx).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}
