package auth

import (
	"net/http"
)

func (h *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	store, _ := h.CookieStore.Get(r, "session")
	store.Values["user_id"] = nil
	store.Values["authenticated"] = false
	store.Save(r, w)

	h.SessionManager.Destroy(r.Context())

	w.WriteHeader(http.StatusOK)
}
