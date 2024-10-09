package books

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Books) Rate(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	bookId := mux.Vars(r)["BookId"]
	rating, err := strconv.ParseFloat(mux.Vars(r)["Rating"], 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rating < 1 || rating > 10 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tx, err := h.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	bookRating := models.BookRating{}
	getSelfBookRatingQuery := squirrel.Select("rating").From(models.BookRating{}.TableName()).Where(squirrel.Eq{"book_id": bookId, "user_id": user.Id})
	err = getSelfBookRatingQuery.RunWith(tx).QueryRow().Scan(&bookRating.Rating)
	if err != nil {
		bookRating.Rating = 0
	}

	// check if user already rated
	updateBookRatingQuery := squirrel.Update(models.BookRating{}.TableName()).Where(squirrel.Eq{"book_id": bookId, "user_id": user.Id}).Set("rating", rating)

	result, err := updateBookRatingQuery.RunWith(tx).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	rows_affected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	if rows_affected == 0 {
		createBookRatingQuery := squirrel.Insert(models.BookRating{}.TableName()).Columns("book_id", "user_id", "rating").Values(bookId, user.Id, rating)
		_, err = createBookRatingQuery.RunWith(tx).Exec()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tx.Rollback()
			return
		}
	}

	getBookQuery := squirrel.Select("*").From(models.Book{}.TableName()).Where(squirrel.Eq{"id": bookId})
	var tmp_rating sql.NullFloat64
	book := models.Book{Rating: new(float32)}
	err = getBookQuery.RunWith(h.DB).QueryRow().Scan(&book.Id, &book.Name, &book.Cover, &book.Author,
		&tmp_rating, &book.RatingCount, &book.Annotation, &book.Price, &book.PageCount, &book.StockCount,
		&book.OrdersCount, &book.PublishedBy, &book.PublishedAt, &book.Created_at, &book.Updated_at)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		tx.Rollback()
		return
	}

	if tmp_rating.Valid {
		rf := float32(tmp_rating.Float64)
		rcf := float32(book.RatingCount)
		orf := float32(bookRating.Rating)
		nrf := float32(rating)
		*book.Rating = rf*rcf - orf + nrf
		if bookRating.Rating == 0 {
			*book.Rating = *book.Rating / float32(book.RatingCount+1)
			book.RatingCount++
		} else {
			*book.Rating = *book.Rating / float32(book.RatingCount)
		}
	} else {
		*book.Rating = float32(bookRating.Rating)
		book.RatingCount++
	}

	query := squirrel.Update(models.Book{}.TableName()).Set("rating", book.Rating).Set("rating_count", book.RatingCount).Where(squirrel.Eq{"id": bookId})

	_, err = query.RunWith(tx).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	tx.Commit()

	w.WriteHeader(http.StatusOK)
}
