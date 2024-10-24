package books

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Books) RequestBook(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)
	f := struct {
		BookName *string `json:"book_name"`
		Comment  *string `json:"comment"`
	}{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(&f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if f.BookName == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if f.Comment == nil {
		f.Comment = new(string)
		*f.Comment = ""
	}

	createRequestQuery := squirrel.Insert(models.Request{}.TableName()).Columns("user_id", "book_name", "comment").Values(user.Id, *f.BookName, *f.Comment)

	_, err = createRequestQuery.RunWith(h.DB).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
