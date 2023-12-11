package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCollectorRepository),
	fx.Provide(NewCollectTaskRepository),
	fx.Provide(NewDeviceRepository),
	fx.Provide(NewDeviceTypeRepository),
	fx.Provide(NewEthernetRepository),
	fx.Provide(NewReportTaskRepository),
	fx.Provide(NewUserRepository),
)
