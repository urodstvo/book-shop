package auth

import (
	"encoding/json"
	"net/http"
)

func (*Auth) Login(w http.ResponseWriter, r *http.Request) {

	f := struct {
		Login    *string `json:"login"`
		Password *string `json:"password"`
	}{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if f.Login == nil || f.Password == nil {
		http.Error(w, "missing login or password", http.StatusBadRequest)
		return
	}

	res := struct {
		Number int `json:"number"`
	}{Number: 42}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
