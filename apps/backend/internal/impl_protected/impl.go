package impl_protected

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_deps"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected/books"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected/carts"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected/orders"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected/payments"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected/users"
	"github.com/urodstvo/book-shop/libs/config"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
)

type Protected struct {
	*books.Books
	*users.Users
	*orders.Orders
	*payments.Payments
	*carts.Carts
}

type Opts struct {
	fx.In

	DB             *sql.DB
	Config         config.Config
	SessionManager *scs.SessionManager

	Logger logger.Logger
}

func New(opts Opts) *Protected {
	d := &impl_deps.Deps{
		DB:             opts.DB,
		Config:         opts.Config,
		SessionManager: opts.SessionManager,
		Logger:         opts.Logger,
	}

	return &Protected{
		Books:    &books.Books{Deps: d},
		Users:    &users.Users{Deps: d},
		Orders:   &orders.Orders{Deps: d},
		Payments: &payments.Payments{Deps: d},
		Carts:    &carts.Carts{Deps: d},
	}
}
