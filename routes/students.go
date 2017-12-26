package routes

func init() {  

    users_routes := app.Party("/students", userMiddleware)

    {
        users_routes.Get("/id", teste)
    }

}