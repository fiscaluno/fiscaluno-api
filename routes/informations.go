package routes

func init() {

    info_routes := app.Party("/info", userMiddleware)

    {
        info_routes.Get("/", teste)
    }

}