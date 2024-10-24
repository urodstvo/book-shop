package session

import (
	"database/sql"
	"encoding/gob"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/urodstvo/book-shop/libs/models"
)

func New(db *sql.DB) *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour * 31
	sessionManager.Cookie.Name = "session_id"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Store = sqlite3store.NewWithCleanupInterval(db, sessionManager.Lifetime)
	// sessionManager.Cookie.SameSite = http.SameSiteNoneMode
	// sessionManager.Cookie.Secure = true
	sessionManager.Cookie.Domain = "localhost"

	gob.Register(models.User{})

	return sessionManager
}
