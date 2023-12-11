package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAccountRoute),
	fx.Provide(NewUserRoute),
	fx.Provide(NewReportTaskRoutes),
	fx.Provide(NewCollectorRoutes),
	fx.Provide(NewEthernetRoutes),
	fx.Provide(NewDeviceTypeRoutes),
	fx.Provide(NewDeviceRoutes),
	fx.Provide(NewCollectTaskRoutes),
	fx.Provide(NewDebugRoutes),
	// 新增路由放在上面
	fx.Provide(NewRoutes),
)
