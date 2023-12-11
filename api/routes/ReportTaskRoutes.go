package routes

import (
	"greatwall/api/controller"
	"greatwall/lib/web"
)

type ReportTaskRoutes struct {
	httpServer           web.HttpServer
	reportTaskController controller.ReportTaskController
}

func NewReportTaskRoutes(
	httpServer web.HttpServer,
	reportTaskController controller.ReportTaskController) ReportTaskRoutes {
	return ReportTaskRoutes{
		httpServer:           httpServer,
		reportTaskController: reportTaskController,
	}
}

func (s *ReportTaskRoutes) Setup() {
	api := s.httpServer.Gin.Group("/api")
	{
		api.POST("/report-task", s.reportTaskController.Add)
		api.PUT("/report-task", s.reportTaskController.Update)
		api.DELETE("/report-task/:name", s.reportTaskController.Delete)
		api.GET("/report-task/:name", s.reportTaskController.Find)
		api.GET("/report-tasks", s.reportTaskController.FindAll)
		api.GET("/report-task/start/:name", s.reportTaskController.Start)
		api.GET("/report-task/stop/:name", s.reportTaskController.Stop)
	}
}
