package routes

import (
	"greatwall/api/controller"
	"greatwall/lib/web"
)

type DeviceRoutes struct {
	httpServer       web.HttpServer
	deviceController controller.DeviceController
}

func NewDeviceRoutes(
	httpServer web.HttpServer,
	deviceController controller.DeviceController) DeviceRoutes {
	return DeviceRoutes{
		httpServer:       httpServer,
		deviceController: deviceController,
	}
}

func (s *DeviceRoutes) Setup() {

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
