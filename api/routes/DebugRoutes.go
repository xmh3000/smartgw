package routes

import (
	"greatwall/api/controller"
	"greatwall/lib/web"
)

type DebugRoutes struct {
	httpServer      web.HttpServer
	debugController controller.DebugController
}

func NewDebugRoutes(
	httpServer web.HttpServer,
	debugController controller.DebugController) DebugRoutes {
	return DebugRoutes{
		httpServer:      httpServer,
		debugController: debugController,
	}
}

func (s *DebugRoutes) Setup() {
	api := s.httpServer.Gin.Group("/api")
	{
		api.POST("/debug/test", s.debugController.Test)
		api.POST("/debug/upgrade", s.debugController.Upgrade)
		api.POST("/debug/system-status", s.debugController.SystemStatus)
		api.POST("/debug/system-reboot", s.debugController.SystemReboot)
		api.POST("/debug/system-ntp", s.debugController.SystemNtp)

		//api.POST("/debug/watchdog-start", s.debugController.WatchdogStart)
		//api.POST("/debug/watchdog-stop", s.debugController.WatchdogStop)
	}
}
