package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	f := struct {
		Name     *string `json:"name"`
		Login    *string `json:"login"`
		Password *string `json:"password"`
	}{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(&f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if f.Login == nil || f.Password == nil || f.Name == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*f.Password), 16)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := models.User{}
	createUserQuery := squirrel.Insert(models.User{}.TableName()).Columns("login", "password", "name").
		Values(*f.Login, string(hashedPassword), *f.Name).Suffix("RETURNING *")

	err = createUserQuery.RunWith(h.DB).QueryRow().Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.SessionManager.Put(r.Context(), "user", user)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successfully!"))
}
