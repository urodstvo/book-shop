package users

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *Users) PutMe(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	f := struct {
		Name     *string `json:"name"`
		Password *string `json:"password"`
	}{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(&f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if f.Name == nil && f.Password == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if f.Name != nil {
		user.Name = *f.Name
	}
	if f.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*f.Password), 16)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)
	}

	updateUserQuery := squirrel.Update(models.User{}.TableName()).SetMap(squirrel.Eq{"name": user.Name, "password": user.Password}).Where(squirrel.Eq{"id": user.Id})
	_, err = updateUserQuery.RunWith(h.DB).Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.SessionManager.Put(r.Context(), "user", user)

	w.WriteHeader(http.StatusOK)
}
