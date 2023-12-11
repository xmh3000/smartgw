package routes

import (
	"smartgw/api/controller"
	"smartgw/api/repository"
	"smartgw/lib/web"
)

type CollectTaskRoute struct {
	httpServer            *web.HttpServer
	collectTaskController *controller.CollectTaskController
	collectTaskRepository repository.CollectTaskRepository
}

func NewCollectTaskRoute(
	httpServer *web.HttpServer,
	CollectTaskController *controller.CollectTaskController,
	collectTaskRepository repository.CollectTaskRepository) *CollectTaskRoute {
	return &CollectTaskRoute{
		httpServer:            httpServer,
		collectTaskController: CollectTaskController,
		collectTaskRepository: collectTaskRepository,
	}
}

func (s *CollectTaskRoute) Setup() {
	s.collectTaskRepository.Migrate()

	api := s.httpServer.Gin.Group("/api")
	{
		api.POST("/collect-task", s.collectTaskController.Add)
		api.PUT("/collect-task", s.collectTaskController.Update)
		api.DELETE("/collect-task/:name", s.collectTaskController.Delete)
		api.GET("/collect-task/:name", s.collectTaskController.Find)
		api.GET("/collect-tasks", s.collectTaskController.FindAll)
		api.GET("/collect-task/start/:name", s.collectTaskController.Start)
		api.GET("/collect-task/stop/:name", s.collectTaskController.Stop)
	}
}
