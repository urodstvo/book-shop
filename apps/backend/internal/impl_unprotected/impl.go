package implunprotected

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
	"guthub.com/urodstvo/book-shop/libs/config"
)

type Opts struct {
	fx.In

	DB             *sql.DB
	Config         config.Config
	SessionManager *scs.SessionManager

	Logger logger.Logger
}
