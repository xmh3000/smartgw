package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAccountRoute),
	fx.Provide(NewUserRoute),
	fx.Provide(NewReportTaskRoute),
	fx.Provide(NewCollectorRoute),
	fx.Provide(NewEthernetRoute),
	fx.Provide(NewDeviceTypeRoute),
	fx.Provide(NewDeviceRoute),
	fx.Provide(NewCollectTaskRoute),
	fx.Provide(NewDebugRoute),
	// 新增路由放在上面
	fx.Provide(NewRoutes),
)
