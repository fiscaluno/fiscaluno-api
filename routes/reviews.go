package routes

import controller "github.com/fiscaluno/fiscaluno-api/controllers/reviews"

func init() {

	reviews_routes := app.Party("/reviews", userMiddleware)

	{
		reviews_routes.Get("/", controller.AllReviews)
		reviews_routes.Get("/general", controller.GeneralReviews)
		reviews_routes.Get("/detailed", controller.DetailedReviews)
		reviews_routes.Get("/institution/{id:int}", controller.ReviewsByInstitution)
	}

}
