package middlewares

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/urodstvo/book-shop/libs/models"
)

func WithAuth(session *scs.SessionManager, next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := session.Get(r.Context(), "user").(models.User)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
