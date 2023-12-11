package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewReportTaskService),
	fx.Provide(NewCollectorService),
	fx.Provide(NewDeviceTypeService),
	fx.Provide(NewDeviceService),
	fx.Provide(NewCollectTaskService),
	fx.Provide(NewEthernetService),
)
