package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	database "github.com/urodstvo/book-shop/apps/backend/internal/db"
	"github.com/urodstvo/book-shop/apps/backend/internal/handlers"
	"github.com/urodstvo/book-shop/apps/backend/internal/session"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
	"guthub.com/urodstvo/book-shop/libs/config"
)

var App = fx.Options(
	fx.Provide(
		database.New,
		session.New,
		config.NewFx,
		logger.NewFx(
			logger.Opts{},
		),
		fx.Annotate(
			func(handlers []handlers.IHandler) *http.ServeMux {
				mux := http.NewServeMux()
				for _, route := range handlers {
					mux.Handle(route.Pattern(), route.Handler())
				}
				return mux
			},
			fx.ParamTags(`group:"handlers"`),
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
