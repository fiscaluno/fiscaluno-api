package routes

func init() {

	infoRoutes := app.Party("/institutions", userMiddleware)

	{
		infoRoutes.Get("/", teste)
	}

}
