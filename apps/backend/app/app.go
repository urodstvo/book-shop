package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	database "github.com/urodstvo/book-shop/apps/backend/internal/db"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_admin"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_unprotected"
	"github.com/urodstvo/book-shop/apps/backend/internal/routes"
	"github.com/urodstvo/book-shop/apps/backend/internal/session"
	"github.com/urodstvo/book-shop/libs/config"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
)

var App = fx.Options(
	fx.Provide(
		config.NewFx,
		logger.NewFx(),
		database.New,
		session.New,
		session.NewAuth,
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
			c := cors.New(
				cors.Options{
					AllowedOrigins:   []string{"localhost:3000", "http://localhost:3000", "frontend:3000", "http://frontend:3000"},
					AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
					AllowedHeaders:   []string{"Content-Type"},
					AllowCredentials: true,
				},
			)

			server := &http.Server{
				Addr:    ":8080",
				Handler: sessionManager.LoadAndSave(c.Handler(mux)),
			}

			lc.Append(
				fx.Hook{
					OnStart: func(_ context.Context) error {
						go func() {
							l.Info("Started", slog.String("port", "8080"))
							// err := server.ListenAndServeTLS("server.crt", "decrypted.key")
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
