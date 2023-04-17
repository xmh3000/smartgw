package bootstrap

import (
	"context"
	"github.com/boltdb/bolt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"smartgw/api/controller"
	"smartgw/api/repository"
	"smartgw/api/routes"
	"smartgw/api/service"
	"smartgw/lib/config"
	"smartgw/lib/database"
	"smartgw/lib/logger"
	"smartgw/lib/web"
)

var Modules = fx.Options(
	config.Module,
	logger.Module,
	database.Module,
	web.Module,

	repository.Module,
	service.Module,
	controller.Module,
	routes.Module,

	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	server *web.HttpServer,
	routes routes.Routes,
	db *bolt.DB,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				zap.S().Info("启动Web服务器...", server.Server.Addr)
				routes.Setup()
				server.Server.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Info("Web服务器正常退出！")
			// 关闭数据库
			db.Close()
			return server.Server.Shutdown(ctx)
		},
	})
}
