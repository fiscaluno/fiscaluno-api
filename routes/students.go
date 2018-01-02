package routes

import controller "github.com/fiscaluno/fiscaluno-api/controllers/students"

func init() {
	users_routes := app.Party("/students", userMiddleware)

	{
		users_routes.Get("/{id:int}", controller.FindById)
		users_routes.Get("/{id:int}/institution", controller.StudentInstitution)
		users_routes.Post("/signup", controller.InsertStudent)
	}
}
