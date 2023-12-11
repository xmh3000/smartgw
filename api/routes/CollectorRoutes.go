package routes

import (
	"smartgw/api/controller"
	"smartgw/lib/web"
)

type CollectorRoutes struct {
	httpServer          *web.HttpServer
	collectorController *controller.CollectorController
}

func NewCollectorRoute(
	httpServer web.HttpServer,
	collectorController controller.CollectorController) CollectorRoutes {
	return CollectorRoutes{
		httpServer:          httpServer,
		collectorController: collectorController,
	}
}

func (s *CollectorRoutes) Setup() {
	api := s.httpServer.Gin.Group("/api")
	{
		api.POST("/collector", s.collectorController.Add)
		api.PUT("/collector", s.collectorController.Update)
		api.DELETE("/collector/:name", s.collectorController.Delete)
		api.GET("/collector/:name", s.collectorController.Find)
		api.GET("/collectors", s.collectorController.FindAll)
	}
}
