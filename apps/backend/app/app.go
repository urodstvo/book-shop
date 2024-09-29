package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/urodstvo/book-shop/apps/backend/internal/session"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
	"guthub.com/urodstvo/book-shop/libs/config"
)

var App = fx.Options(
	fx.Provide(
		session.New,
		config.NewFx,
		logger.NewFx(
			logger.Opts{},
		),
	),
	fx.Invoke(
		func(
			mux *http.ServeMux,
			l logger.Logger,
			lc fx.Lifecycle,
		) error {
			server := &http.Server{
				Addr: "0.0.0.0:8000",
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
