package routes

import (
	"greatwall/api/controller"
	"greatwall/api/repository"
	"greatwall/lib/web"
)

type CollectTaskRoutes struct {
	httpServer            web.HttpServer
	collectTaskController controller.CollectTaskController
	collectTaskRepository repository.CollectTaskRepository
}

func NewCollectTaskRoutes(
	httpServer web.HttpServer,
	CollectTaskController controller.CollectTaskController,
	collectTaskRepository repository.CollectTaskRepository) CollectTaskRoutes {
	return CollectTaskRoutes{
		httpServer:            httpServer,
		collectTaskController: CollectTaskController,
		collectTaskRepository: collectTaskRepository,
	}
}

func (s *CollectTaskRoutes) Setup() {
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
