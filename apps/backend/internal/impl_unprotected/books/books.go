package books

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_deps"
	"github.com/urodstvo/book-shop/libs/models"
)

type Books struct {
	*impl_deps.Deps
}

type bookResponse struct {
	*models.Book
	UserRating *int           `json:"user_rating"`
	Genres     []models.Genre `json:"genres"`
}

func (h *Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	user, ok := h.SessionManager.Get(r.Context(), "user").(models.User)

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	author := r.URL.Query().Get("author")
	name := r.URL.Query().Get("name")
	published_by := r.URL.Query().Get("pubby")
	published_after := r.URL.Query().Get("puba")
	published_before := r.URL.Query().Get("pubb")
	rating_lower := r.URL.Query().Get("rb")
	rating_bigger := r.URL.Query().Get("ra")
	price_order := r.URL.Query().Get("po")
	genres := r.URL.Query()["genres"]

	if page == "" || limit == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := squirrel.Select("*").From(models.Book{}.TableName())

	if author != "" {
		query = query.Where(squirrel.Like{"author": "%" + author + "%"})
	}

	if name != "" {
		query = query.Where(squirrel.Like{"name": "%" + name + "%"})
	}

	if published_by != "" {
		query = query.Where(squirrel.Eq{"published_by": "%" + published_by + "%"})
	}

	if published_after != "" {
		query = query.Where(squirrel.GtOrEq{"published_at": published_after})
	}

	if published_before != "" {
		query = query.Where(squirrel.LtOrEq{"published_at": published_before})
	}

	if rating_lower != "" {
		query = query.Where(squirrel.LtOrEq{"rating": rating_lower})
	}

	if rating_bigger != "" {
		query = query.Where(squirrel.GtOrEq{"rating": rating_bigger})
	}

	if len(genres) > 0 {
		book_ids := []int{}
		booksWithGenreQuery := squirrel.Select("book_id").Distinct().From(models.BookGenre{}.TableName()).
			Where(squirrel.Eq{"genre_id": genres}).GroupBy("book_id").Having(squirrel.Expr("COUNT(genre_id) = ?", len(genres)))
		rows, err := booksWithGenreQuery.RunWith(h.DB).Query()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			bookGenre := models.BookGenre{}
			err = rows.Scan(&bookGenre.BookId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			book_ids = append(book_ids, bookGenre.BookId)
		}
		query = query.Where(squirrel.Eq{"id": book_ids})
	}

	if price_order != "" {
		if price_order == "asc" {
			query = query.OrderBy("price ASC")
		} else if price_order == "desc" {
			query = query.OrderBy("price DESC")
		}
	}

	query = query.Limit(uint64(limitNum)).Offset(uint64((pageNum - 1) * limitNum))

	books := []bookResponse{}
	rows, err := query.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var tmp_rating sql.NullFloat64
		book := bookResponse{Book: &models.Book{}, Genres: []models.Genre{}}
		err = rows.Scan(&book.Book.Id, &book.Book.Name, &book.Book.Cover, &book.Book.Author,
			&tmp_rating, &book.Book.RatingCount, &book.Book.Annotation, &book.Book.Price, &book.Book.PageCount, &book.Book.StockCount,
			&book.Book.OrdersCount, &book.Book.PublishedBy, &book.Book.PublishedAt, &book.Book.Created_at, &book.Book.Updated_at)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		getGenresQuery := squirrel.Select("g.id", "g.name").From(models.BookGenre{}.TableName() + " b").
			Join(models.Genre{}.TableName() + " g on b.genre_id = g.id").Where(squirrel.Eq{"b.book_id": book.Book.Id})

		gRows, err := getGenresQuery.RunWith(h.DB).Query()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for gRows.Next() {
			genre := models.Genre{}
			err = gRows.Scan(&genre.Id, &genre.Name)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			book.Genres = append(book.Genres, genre)
		}

		if tmp_rating.Valid {
			book.Rating = new(float32)
			*book.Rating = float32(tmp_rating.Float64)
		} else {
			book.Rating = nil
		}

		var userRating *int = nil
		if ok {
			getRatingQuery := squirrel.Select("rating").From(models.BookRating{}.TableName()).Where(squirrel.Eq{"user_id": user.Id, "book_id": book.Id})
			_ = getRatingQuery.RunWith(h.DB).QueryRow().Scan(&userRating)
		}

		book.UserRating = userRating

		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&books)
}
