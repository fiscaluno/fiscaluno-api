package routes

import controller "github.com/fiscaluno/fiscaluno-api/controllers/institutions"

func init() {

	info_routes := app.Party("/institutions", userMiddleware)

	{
		info_routes.Get("/", controller.AllInstitutions)
	}

}
