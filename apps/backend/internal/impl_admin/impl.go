package impl_admin

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
	"guthub.com/urodstvo/book-shop/libs/config"
)

type Admin struct {
}

type Opts struct {
	fx.In

	DB             *sql.DB
	Config         config.Config
	SessionManager *scs.SessionManager

	Logger logger.Logger
}

func New(opts Opts) *Admin {
	// d := &impl_deps.Deps{
	// 	DB:             opts.DB,
	// 	Config:         opts.Config,
	// 	SessionManager: opts.SessionManager,
	// 	Logger:         opts.Logger,
	// }

	return &Admin{}
}
