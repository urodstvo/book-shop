package orders

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

type response struct {
	Orders int       `json:"orders"`
	Date   time.Time `json:"date"`
}

func (h *Orders) GetDynamic(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	if start == "" || end == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
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

	getOrdersQuery := squirrel.Select("COUNT(*)").From(models.Order{}.TableName()).Where("NOT(status = ?)", "cancelled")

	res := []response{}
	for startTime.Before(endTime.AddDate(0, 0, 1)) {
		getOrdersOnDayQuery := getOrdersQuery.Where("DATE(created_at) = ?", startTime.Format("2006-01-02"))

		rows, err := getOrdersOnDayQuery.RunWith(h.DB).Query()
		if err != nil {

			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res = append(res, response{Orders: 0, Date: startTime})

		for rows.Next() {
			var count float64
			err = rows.Scan(&count)
			if err != nil {
				h.Logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			res[len(res)-1].Orders = int(count)
		}

		startTime = startTime.AddDate(0, 0, 1)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
