package routes

func init() {

	usersRoutes := app.Party("/students", userMiddleware)

	{
		userRoutes.Get("/id", teste)
	}

}
