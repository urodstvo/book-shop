package books

import (
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

func (h *Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	author := r.URL.Query().Get("author")
	name := r.URL.Query().Get("name")
	published_after := r.URL.Query().Get("puba")
	published_before := r.URL.Query().Get("pubb")
	rating_lower := r.URL.Query().Get("rb")
	rating_bigger := r.URL.Query().Get("ra")
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
		booksWithGenreQuery := squirrel.Select("book_id").Distinct().From(models.BookGenre{}.TableName()).Where(squirrel.Eq{"genre_id": genres})
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

	query = query.Limit(uint64(limitNum)).Offset(uint64((pageNum - 1) * limitNum))

	books := []models.Book{}
	rows, err := query.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.Id, &book.Name, &book.Cover, &book.Author,
			&book.Rating, &book.RatingCount, &book.Annotation, &book.PageCount, &book.StockCount,
			&book.OrdersCount, &book.PublishedBy, &book.PublishedAt, &book.Created_at, &book.Updated_at)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
