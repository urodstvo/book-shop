package main

import (
	"github.com/urodstvo/book-shop/apps/backend/app"
	"github.com/urodstvo/book-shop/libs/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		logger.FxDiOnlyErrors,
		app.App,
	).Run()
}
