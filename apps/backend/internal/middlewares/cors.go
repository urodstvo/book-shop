package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func WithCors(base http.Handler) http.Handler {
	corsWrapper := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"Content-Type"},
			AllowCredentials: true,
		},
	)
	handler := corsWrapper.Handler(base)

	return handler
}
