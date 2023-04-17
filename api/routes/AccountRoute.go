package routes

import (
	"smartgw/api/controller"
	"smartgw/lib/web"
	"smartgw/lib/web/middleware"
)

type AccountRoute struct {
	server            *web.HttpServer
	accountController *controller.AccountController
}

var _ Route = (*AccountRoute)(nil)

func NewAccountRoute(
	server *web.HttpServer,
	accountController *controller.AccountController,
) *AccountRoute {
	return &AccountRoute{
		server:            server,
		accountController: accountController,
	}
}

func (a *AccountRoute) Setup() {
	router := a.server.Gin
	router.POST("/login", a.accountController.Login)

	router.Use(middleware.JWTAuth())
}
