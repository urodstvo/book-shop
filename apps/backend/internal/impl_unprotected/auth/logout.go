package auth

import (
	"net/http"
)

func (h *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	h.SessionManager.Destroy(r.Context())

	w.WriteHeader(http.StatusOK)
}
