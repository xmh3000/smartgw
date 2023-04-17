package routes

import (
	"smartgw/api/controller"
	"smartgw/api/repository"
	"smartgw/lib/web"
)

type UserRoute struct {
	*web.HttpServer
	repository.UserRepository
	*controller.UserController
}

var _ Route = (*UserRoute)(nil)

func NewUserRoute(
	server *web.HttpServer,
	userRepository repository.UserRepository,
	userController *controller.UserController,
) *UserRoute {
	return &UserRoute{
		server,
		userRepository,
		userController,
	}
}

func (u *UserRoute) Setup() {
	u.Migrate()

	api := u.Gin.Group("/api")
	{
		api.GET("/user/:name", u.UserController.Find)
	}
}
