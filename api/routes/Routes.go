package routes

type (
	Routes []Route
	Route  interface {
		Setup()
	}
)

func NewRoutes(
	accountRoute *AccountRoute,
	userRoute *UserRoute,
	reportTaskRoute *ReportTaskRoute,
	collectorRoute *CollectorRoute,
	ethernetRoute *EthernetRoute,
	deviceTypeRoute *DeviceTypeRoute,
	deviceRoute *DeviceRoute,
	collectTaskRoute *CollectTaskRoute,
	debugRoute *DebugRoute,
) Routes {
	return Routes{
		accountRoute,
		userRoute,
		reportTaskRoute,
		collectorRoute,
		ethernetRoute,
		deviceTypeRoute,
		deviceRoute,
		collectTaskRoute,
		debugRoute,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
