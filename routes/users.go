package routes

func init() {  

    users_routes := app.Party("/users", userMiddleware)

    {
        users_routes.Get("/id", teste)
    }

}