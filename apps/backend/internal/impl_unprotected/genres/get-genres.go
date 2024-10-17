package genres

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Genres) GetGenres(w http.ResponseWriter, r *http.Request) {

	query := squirrel.Select("*").From(models.Genre{}.TableName())

	genres := []models.Genre{}

	rows, err := query.RunWith(h.DB).Query()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		genre := models.Genre{}
		err = rows.Scan(&genre.Id, &genre.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		genres = append(genres, genre)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genres)

}
