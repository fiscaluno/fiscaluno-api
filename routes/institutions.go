package routes

import controller "github.com/fiscaluno/fiscaluno-api/controllers/institutions"

func init() {

	institutions_routes := app.Party("/institutions", userMiddleware)

	{
		institutions_routes.Get("/", controller.AllInstitutions)
		institutions_routes.Get("/{id:int}", controller.InstitutionById)
        institutions_routes.Get("/{id:int}/reviews", controller.InstitutionReviews)
        institutions_routes.Get("/{id:int}/reviews/general", controller.InstitutionGeneralReviews)
		institutions_routes.Get("/{id:int}/reviews/detailed", controller.InstitutionDetailedReviews)
	}

}
