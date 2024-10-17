package books

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

type response struct {
	Rating float64   `json:"rating"`
	Date   time.Time `json:"date"`
}

func (h *Books) GetRating(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	author := r.URL.Query().Get("author")
	name := r.URL.Query().Get("name")
	genres := r.URL.Query()["genres"]

	if start == "" || end == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getBookRatingsQuery := squirrel.Select("AVG(r.rating), DATE(r.updated_at) as rating_date").From(models.BookRating{}.TableName() + " r").GroupBy("rating_date")

	if author != "" || name != "" {
		getBookRatingsQuery = getBookRatingsQuery.Join(models.Book{}.TableName() + " b on b.id = r.book_id")
		if author != "" {
			getBookRatingsQuery = getBookRatingsQuery.Where(squirrel.Eq{"b.author": author})
		}
		if name != "" {
			getBookRatingsQuery = getBookRatingsQuery.Where(squirrel.Like{"b.name": name})
		}
	}

	if len(genres) > 0 {
		book_ids := []int{}
		booksWithGenreQuery := squirrel.Select("book_id").Distinct().From(models.BookGenre{}.TableName()).
			Where(squirrel.Eq{"genre_id": genres}).GroupBy("book_id").Having(squirrel.Expr("COUNT(genre_id) = ?", len(genres)))
		rows, err := booksWithGenreQuery.RunWith(h.DB).Query()
		if err != nil {
			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			bookGenre := models.BookGenre{}
			err = rows.Scan(&bookGenre.BookId)
			if err != nil {

				h.Logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			book_ids = append(book_ids, bookGenre.BookId)
		}
		getBookRatingsQuery = getBookRatingsQuery.Where(squirrel.Eq{"r.book_id": book_ids})
	}

	startTime, err := time.Parse("2006-01-02", start[0:10])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	endTime, err := time.Parse("2006-01-02", end[0:10])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := []response{}
	for startTime.Before(endTime.AddDate(0, 0, 1)) {
		getRatingsOnDayQuery := getBookRatingsQuery.Where("DATE(r.updated_at) <= ?", startTime.Format("2006-01-02"))

		// sql, args, err := getRatingsOnDayQuery.ToSql()
		// if err != nil {
		// 	h.Logger.Error(err.Error())
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		//
		// h.Logger.Info(sql, args)

		rows, err := getRatingsOnDayQuery.RunWith(h.DB).Query()
		if err != nil {

			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res = append(res, response{Rating: 0, Date: startTime})

		for rows.Next() {
			var avg_rating float64
			var stringDate string
			err = rows.Scan(&avg_rating, &stringDate)
			if err != nil {
				h.Logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			res[len(res)-1].Rating = avg_rating
		}

		startTime = startTime.AddDate(0, 0, 1)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
