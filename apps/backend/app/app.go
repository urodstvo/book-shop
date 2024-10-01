package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
	database "github.com/urodstvo/book-shop/apps/backend/internal/db"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_admin"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_unprotected"
	"github.com/urodstvo/book-shop/apps/backend/internal/routes"
	"github.com/urodstvo/book-shop/apps/backend/internal/session"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
	"guthub.com/urodstvo/book-shop/libs/config"
)

var App = fx.Options(
	fx.Provide(
		database.New,
		session.New,
		session.NewAuth,
		config.NewFx,
		logger.NewFx(
			logger.Opts{},
		),
		impl_unprotected.New,
		impl_protected.New,
		impl_admin.New,
		routes.New,
	),

	fx.Invoke(
		func(
			mux *mux.Router,
			sessionManager *scs.SessionManager,
			l logger.Logger,
			lc fx.Lifecycle,
		) error {
			server := &http.Server{
				Addr:    "0.0.0.0:8000",
				Handler: sessionManager.LoadAndSave(mux),
			}

			lc.Append(
				fx.Hook{
					OnStart: func(_ context.Context) error {
						go func() {
							l.Info("Started", slog.String("port", "8000"))
							err := server.ListenAndServe()
							if err != nil && !errors.Is(err, http.ErrServerClosed) {
								panic(err)
							}
						}()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						return server.Shutdown(ctx)
					},
				},
			)

			return nil
		},
	),
)
