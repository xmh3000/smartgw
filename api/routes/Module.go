package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAccountRoute),
	fx.Provide(NewUserRoute),
	fx.Provide(NewRoutes),
)
