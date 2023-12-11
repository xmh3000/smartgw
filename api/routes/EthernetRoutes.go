package routes

import (
	"greatwall/api/controller"
	"greatwall/api/repository"
	"greatwall/lib/web"
)

type EthernetRoutes struct {
	httpServer         web.HttpServer
	ethernetController controller.EthernetController
	ethernetRepository repository.EthernetRepository
}

func NewEthernetRoutes(
	httpServer web.HttpServer,
	ethernetController controller.EthernetController,
	ethernetRepository repository.EthernetRepository,
) EthernetRoutes {
	return EthernetRoutes{
		httpServer:         httpServer,
		ethernetController: ethernetController,
		ethernetRepository: ethernetRepository,
	}
}

func (s *EthernetRoutes) Setup() {
	// 不在软件里设置网卡ip了，以后都放到rc.local中
	// s.ethernetRepository.Migrate()

	api := s.httpServer.Gin.Group("/api")
	{
		//api.POST("/ethernet", s.ethernetController.Add)
		api.PUT("/ethernet", s.ethernetController.Update)
		api.DELETE("/ethernet/:name", s.ethernetController.Delete)
		api.GET("/ethernet/:name", s.ethernetController.Find)
		api.GET("/ethernets", s.ethernetController.FindAll)
	}
}
