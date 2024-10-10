package books

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

type Recomendation struct {
	*models.Book
	Genres []models.Genre `json:"genres"`
}

func (h *Books) Recomendations(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	getAllRatedBooksByUserQuery := squirrel.Select("book_id").From(models.BookRating{}.TableName()).Where(squirrel.Eq{"user_id": user.Id})
	getAllOrderedBooksByUserQuery := squirrel.Select("book_id").From(models.OrderBook{}.TableName() + " ob").Join(models.Order{}.TableName() + " o on o.id = ob.order_id").Where(squirrel.Eq{"o.user_id": user.Id})

	rows, err := getAllRatedBooksByUserQuery.RunWith(h.DB).Query()
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	books := []int{}

	for rows.Next() {
		var book_id int

		err = rows.Scan(&book_id)
		if err != nil {

			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		books = append(books, book_id)
	}

	rows, err = getAllOrderedBooksByUserQuery.RunWith(h.DB).Query()
	if err != nil {

		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var book_id int

		err = rows.Scan(&book_id)
		if err != nil {

			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		books = append(books, book_id)
	}

	getAllNotRatedAndNotOrderedBooksQuery := squirrel.Select("*").From(models.Book{}.TableName()).Where(squirrel.NotEq{"id": books})

	rows, err = getAllNotRatedAndNotOrderedBooksQuery.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	recomentations := []Recomendation{}

	for rows.Next() {
		var tmp_rating sql.NullFloat64
		book := &models.Book{Rating: new(float32)}

		err = rows.Scan(&book.Id, &book.Name, &book.Cover, &book.Author, &tmp_rating, &book.RatingCount, &book.Annotation, &book.Price, &book.PageCount,
			&book.StockCount, &book.OrdersCount, &book.PublishedBy, &book.PublishedAt, &book.Created_at, &book.Updated_at)

		if err != nil {

			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if tmp_rating.Valid {
			*book.Rating = float32(tmp_rating.Float64)
		} else {
			book.Rating = nil
		}

		getGenresQuery := squirrel.Select("g.id", "g.name").From(models.BookGenre{}.TableName() + " b").Join(models.Genre{}.TableName() + " g on b.genre_id = g.id").Where(squirrel.Eq{"b.book_id": book.Id})

		genreRows, err := getGenresQuery.RunWith(h.DB).Query()
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		genres := []models.Genre{}

		for genreRows.Next() {
			genre := models.Genre{}

			err = genreRows.Scan(&genre.Id, &genre.Name)
			if err != nil {

				h.Logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			genres = append(genres, genre)
		}

		recomentations = append(recomentations, Recomendation{book, genres})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recomentations)
}
