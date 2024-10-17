package books

import (
	"net/http"
	"text/template"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Books) BookPreview(w http.ResponseWriter, r *http.Request) {
	book_id := mux.Vars(r)["BookId"]

	getBookContentQuery := squirrel.Select("*").From(models.BookPage{}.TableName()).Where(squirrel.Eq{"book_id": book_id})

	content := models.BookPage{}

	err := getBookContentQuery.RunWith(h.DB).QueryRow().Scan(&content.BookId, &content.Content)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl, err := template.New("book_preview").Parse(`{{.Content}}`)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, content)
}
