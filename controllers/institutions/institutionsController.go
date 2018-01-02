package institutions

import (
	institution "github.com/fiscaluno/fiscaluno-api/models/institution"
	"github.com/kataras/iris"
	"reflect"
)

// List of all institutions - GET Method
func AllInstitutions(ctx iris.Context) {
	ctx.JSON(institution.All())
}

// Institution informations - GET Method
func InstitutionById(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.JSON(err)
	}

	// Checks if student exists at database
	if reflect.DeepEqual(institution.Institution{}, *institution.Find(id)) {
		ctx.StatusCode(404)
	} else {
		ctx.JSON(institution.Find(id))
	}
}

// All reviews from institution
func InstitutionReviews(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.JSON(err)
	}

	ctx.JSON(id)
}

// All general reviews from institution
func InstitutionGeneralReviews(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.JSON(err)
	}

	ctx.JSON(institution.Find(id).GeneralReviews())
}

// All detailed reviews from institution
func InstitutionDetailedReviews(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")

	if err != nil {
		ctx.JSON(err)
	}

	ctx.JSON(institution.Find(id).DetailedReviews())
}
