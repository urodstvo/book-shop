package impl_unprotected

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/sessions"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_deps"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_unprotected/auth"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_unprotected/books"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_unprotected/hello_world"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
	"guthub.com/urodstvo/book-shop/libs/config"
)

type UnProtected struct {
	*auth.Auth
	*hello_world.HelloWorld
	*books.Books
}

type Opts struct {
	fx.In

	DB             *sql.DB
	Config         config.Config
	SessionManager *scs.SessionManager
	CookieStore    *sessions.CookieStore

	Logger logger.Logger
}

func New(opts Opts) *UnProtected {
	d := &impl_deps.Deps{
		DB:             opts.DB,
		Config:         opts.Config,
		SessionManager: opts.SessionManager,
		CookieStore:    opts.CookieStore,
		Logger:         opts.Logger,
	}

	return &UnProtected{
		Auth:       &auth.Auth{Deps: d},
		HelloWorld: &hello_world.HelloWorld{Deps: d},
		Books:      &books.Books{Deps: d},
	}
}
