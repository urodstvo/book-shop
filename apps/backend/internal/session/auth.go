package session

import "github.com/gorilla/sessions"

func NewAuth() *sessions.CookieStore {
	s := sessions.NewCookieStore([]byte("secret-key"))
	s.Options.HttpOnly = true

	return s
}
