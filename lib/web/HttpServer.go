package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartgw/lib/config"
	"smartgw/lib/logger"
)

type HttpServer struct {
	Gin    *gin.Engine
	Server *http.Server
}

func NewHttpServer(config *config.Config) *HttpServer {
	gin.SetMode(config.Server.Mode)
	engine := gin.Default()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))
	engine.SetTrustedProxies([]string{config.Server.Address})

	engine.Static("/static", "./webroot/static")
	engine.StaticFile("/", "./webroot/index.html")
	engine.StaticFile("/favicon.ico", "./webroot/favicon.ico")

	return &HttpServer{
		Gin:    engine,
		Server: &http.Server{Addr: config.Server.Port, Handler: engine},
	}
}
