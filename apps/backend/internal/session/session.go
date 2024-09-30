package session

import (
	"database/sql"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
)

func New(db *sql.DB) *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour * 31
	sessionManager.Cookie.Name = "session_id"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Store = sqlite3store.NewWithCleanupInterval(db, sessionManager.Lifetime)

	return sessionManager
}
