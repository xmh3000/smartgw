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
	"smartgw/lib/collect"
	"smartgw/lib/config"
	"smartgw/lib/database"
	"smartgw/lib/io"
	"smartgw/lib/logger"
	"smartgw/lib/mqtt"
	"smartgw/lib/task"
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
	collect.Module,
	task.Module,

	// 获取系统运行
	fx.Invoke(io.GetStartTime),
	// 初始化配置文件、日志、数据库等
	//fx.Invoke(logger.InitLogger),
	// 启动ntp服务
	fx.Invoke(io.InitNtp),
	// 启动看门狗
	//fx.Invoke(watchdog.InitWatchDog),

	//fx.Invoke(led.InitLed),

	fx.Invoke(mqtt.InitMqttClient),
	// 启动采集接口服务器
	fx.Invoke(collect.InitCollectorServer),
	// 启动采集任务服务器
	fx.Invoke(task.InitCollectTaskServer),
	// 启动上报任务服务器
	fx.Invoke(task.InitReportTaskServer),
	// 启动 mqttClient、web
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
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				zap.S().Info("启动 Mqtt 客户端...")
				mqtt.MClient.Connect()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Info("停止 Mqtt 客户端...")

			mqtt.MClient.Discount()
			return nil
		},
	})

}
