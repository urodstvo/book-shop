package books

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Books) GetByBookId(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["BookId"]

	query := squirrel.Select("*").From(models.Book{}.TableName()).Where(squirrel.Eq{"id": bookId})
	book := &models.Book{}
	err := query.RunWith(h.DB).QueryRow().Scan(&book.Id, &book.Name, &book.Cover, &book.Author,
		&book.Rating, &book.RatingCount, &book.Annotation, &book.PageCount, &book.StockCount,
		&book.OrdersCount, &book.PublishedBy, &book.PublishedAt, &book.Created_at, &book.Updated_at)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
