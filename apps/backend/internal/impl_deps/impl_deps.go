package impl_deps

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/sessions"
	"github.com/urodstvo/book-shop/libs/logger"
	"guthub.com/urodstvo/book-shop/libs/config"
)

type Deps struct {
	Config         config.Config
	DB             *sql.DB
	SessionManager *scs.SessionManager
	CookieStore    *sessions.CookieStore
	Logger         logger.Logger
}
