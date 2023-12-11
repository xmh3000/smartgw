package routes

import (
	"smartgw/api/controller"
	"smartgw/api/repository"
	"smartgw/lib/web"
)

type EthernetRoute struct {
	httpServer         *web.HttpServer
	ethernetController *controller.EthernetController
	ethernetRepository repository.EthernetRepository
}

func NewEthernetRoute(
	httpServer *web.HttpServer,
	ethernetController *controller.EthernetController,
	ethernetRepository repository.EthernetRepository,
) *EthernetRoute {
	return &EthernetRoute{
		httpServer:         httpServer,
		ethernetController: ethernetController,
		ethernetRepository: ethernetRepository,
	}
}

func (s *EthernetRoute) Setup() {
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
