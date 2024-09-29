package impl_deps

import (
	"github.com/alexedwards/scs/v2"
	"github.com/urodstvo/book-shop/apps/backend/internal/repo"
	"github.com/urodstvo/book-shop/libs/logger"
	"guthub.com/urodstvo/book-shop/libs/config"
)

type Deps struct {
	Config         config.Config
	DB             *repo.Database
	SessionManager *scs.SessionManager
	Logger         logger.Logger
}
