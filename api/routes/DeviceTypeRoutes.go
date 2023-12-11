package routes

import (
	"greatwall/api/controller"
	"greatwall/lib/web"
)

type DeviceTypeRoutes struct {
	httpServer           web.HttpServer
	deviceTypeController controller.DeviceTypeController
}

func NewDeviceTypeRoutes(
	httpServer web.HttpServer,
	deviceTypeController controller.DeviceTypeController) DeviceTypeRoutes {
	return DeviceTypeRoutes{
		httpServer:           httpServer,
		deviceTypeController: deviceTypeController,
	}
}

func (s *DeviceTypeRoutes) Setup() {
	api := s.httpServer.Gin.Group("/api")
	{
		api.POST("/device-type", s.deviceTypeController.Add)
		api.PUT("/device-type", s.deviceTypeController.Update)
		api.DELETE("/device-type/:name", s.deviceTypeController.Delete)
		api.GET("/device-type/:name", s.deviceTypeController.Find)
		api.GET("/device-types", s.deviceTypeController.FindAll)
		api.POST("/device-type/upload/:name", s.deviceTypeController.Upload)

		// id为设备类型id，propertyId为属性id
		api.POST("/device-property/:name", s.deviceTypeController.AddProperties)
		api.PUT("/device-property/:name/:propertyid", s.deviceTypeController.UpdateProperties)
		api.DELETE("/device-property/:name/:propertyid", s.deviceTypeController.DeleteProperties)
		api.GET("/device-property/:name/:propertyid", s.deviceTypeController.FindProperty)
		api.GET("/device-property/:name", s.deviceTypeController.FindAllProperties)
	}
}
