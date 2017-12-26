package routes

func init() {

    info_routes := app.Party("/institutions", userMiddleware)

    {
        info_routes.Get("/", teste)
    }

}