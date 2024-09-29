package session

import (
	"time"

	"github.com/alexedwards/scs/v2"
)

func New() *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour * 31
	sessionManager.Cookie.Name = "session_id"
	sessionManager.Cookie.HttpOnly = true

	return sessionManager
}
