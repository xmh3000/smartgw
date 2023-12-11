package routes

import (
	"smartgw/api/controller"
	"smartgw/lib/web"
)

type DeviceRoute struct {
	httpServer       *web.HttpServer
	deviceController *controller.DeviceController
}

func NewDeviceRoute(
	httpServer *web.HttpServer,
	deviceController *controller.DeviceController) *DeviceRoute {
	return &DeviceRoute{
		httpServer:       httpServer,
		deviceController: deviceController,
	}
}

func (s *DeviceRoute) Setup() {

	api := s.httpServer.Gin.Group("/api")
	{
		api.POST("/device", s.deviceController.Add)
		api.PUT("/device", s.deviceController.Update)
		api.DELETE("/device/:name", s.deviceController.Delete)
		api.GET("/device/:name", s.deviceController.Find)
		api.GET("/devices", s.deviceController.FindAll)
		api.POST("/devices/import", s.deviceController.Import)
		api.GET("/devices/export", s.deviceController.Export)
	}
}
