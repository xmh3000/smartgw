package web

import (
	"context"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"smartgw/lib/config"
	"smartgw/lib/logger"
)

func NewGin(lifecycle fx.Lifecycle, config *config.Config, db *bolt.DB) *gin.Engine {
	gin.SetMode(config.Server.Mode)
	engine := gin.Default()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))
	engine.SetTrustedProxies([]string{config.Server.Address})

	engine.Static("/static", "./webroot/static")
	engine.StaticFile("/", "./webroot/index.html")
	engine.StaticFile("/favicon.ico", "./webroot/favicon.ico")

	server := &http.Server{Addr: config.Server.Port, Handler: engine}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				zap.S().Info("启动Web服务器...", server.Addr)
				server.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Info("Web服务器正常退出！")
			// 关闭数据库
			db.Close()
			return server.Shutdown(ctx)
		},
	})

	return engine
}
