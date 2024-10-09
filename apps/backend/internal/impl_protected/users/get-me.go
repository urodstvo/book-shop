package users

import (
	"encoding/json"
	"net/http"

	"github.com/urodstvo/book-shop/libs/models"
)

func (h *Users) GetMe(w http.ResponseWriter, r *http.Request) {
	user := h.SessionManager.Get(r.Context(), "user").(models.User)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
