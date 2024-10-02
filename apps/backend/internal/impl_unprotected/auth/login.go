package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/urodstvo/book-shop/libs/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	f := struct {
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

	if f.Login == nil || f.Password == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := models.User{}
	getUserQuery := squirrel.Select("*").From(models.User{}.TableName()).
		Where("login = ?", *f.Login)

	err = getUserQuery.RunWith(h.DB).QueryRow().Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Role)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*f.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	store, _ := h.CookieStore.Get(r, "session")
	store.Values["user_id"] = user.Id
	store.Values["authenticated"] = true
	store.Values["role"] = user.Role
	store.Save(r, w)

	h.SessionManager.Put(r.Context(), "user_id", user.Id)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successfully!"))
}
