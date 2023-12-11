package routes

import (
	"smartgw/api/controller"
	"smartgw/lib/web"
)

type CollectorRoute struct {
	httpServer          *web.HttpServer
	collectorController *controller.CollectorController
}

func NewCollectorRoute(
	httpServer *web.HttpServer,
	collectorController *controller.CollectorController) *CollectorRoute {
	return &CollectorRoute{
		httpServer:          httpServer,
		collectorController: collectorController,
	}
}

func (s *CollectorRoute) Setup() {
	api := s.httpServer.Gin.Group("/api")
	{
		api.POST("/collector", s.collectorController.Add)
		api.PUT("/collector", s.collectorController.Update)
		api.DELETE("/collector/:name", s.collectorController.Delete)
		api.GET("/collector/:name", s.collectorController.Find)
		api.GET("/collectors", s.collectorController.FindAll)
	}
}
