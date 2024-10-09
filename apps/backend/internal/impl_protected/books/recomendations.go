package books

import "net/http"

func (h *Books) Recomendations(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
