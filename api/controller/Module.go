package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAccountController),
	fx.Provide(NewUserController),
	fx.Provide(NewReportTaskController),
	fx.Provide(NewCollectorController),
	fx.Provide(NewDeviceTypeController),
	fx.Provide(NewDeviceController),
	fx.Provide(NewCollectTaskController),
	fx.Provide(NewDebugController),
	fx.Provide(NewEthernetController),
)
