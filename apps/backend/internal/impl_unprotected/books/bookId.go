package books

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Books) GetByBookId(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["BookId"]
	var userRating *int = nil
	user, ok := h.SessionManager.Get(r.Context(), "user").(models.User)
	if ok {
		query := squirrel.Select("rating").From(models.BookRating{}.TableName()).Where(squirrel.Eq{"user_id": user.Id, "book_id": bookId})
		_ = query.RunWith(h.DB).QueryRow().Scan(&userRating)
	}

	query := squirrel.Select("*").From(models.Book{}.TableName()).Where(squirrel.Eq{"id": bookId})
	var tmp_rating sql.NullFloat64
	book := &models.Book{}
	err := query.RunWith(h.DB).QueryRow().Scan(&book.Id, &book.Name, &book.Cover, &book.Author,
		&tmp_rating, &book.RatingCount, &book.Annotation, &book.Price, &book.PageCount, &book.StockCount,
		&book.OrdersCount, &book.PublishedBy, &book.PublishedAt, &book.Created_at, &book.Updated_at)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if tmp_rating.Valid {
		book.Rating = new(float32)
		*book.Rating = float32(tmp_rating.Float64)
	} else {
		book.Rating = nil
	}

	getGenresQuery := squirrel.Select("g.id", "g.name").From(models.BookGenre{}.TableName() + " b").
		Join(models.Genre{}.TableName() + " g on b.genre_id = g.id").Where(squirrel.Eq{"b.book_id": book.Id})

	gRows, err := getGenresQuery.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := bookResponse{
		book,
		userRating,
		[]models.Genre{},
	}

	for gRows.Next() {
		genre := models.Genre{}
		err = gRows.Scan(&genre.Id, &genre.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Genres = append(res.Genres, genre)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
