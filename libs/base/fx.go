package base

import "go.uber.org/fx"

func CreateBaseApp() fx.Option {
	return fx.Options(
		fx.Provide(),
	)
}
