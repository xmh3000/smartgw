package routes

type (
	Routes []Route
	Route  interface {
		Setup()
	}
)

func NewRoutes(
	userRoute *UserRoute,
	accountRoute *AccountRoute,
) Routes {
	return Routes{
		userRoute,
		accountRoute,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
