package reviews

import (
	"fmt"
	"github.com/fiscaluno/fiscaluno-api/models/detailedReview"
	"github.com/fiscaluno/fiscaluno-api/models/generalReview"
	"github.com/fiscaluno/fiscaluno-api/models/institution"
	"github.com/kataras/iris"
)

func AllReviews(ctx iris.Context) {
	generalReviews := generalReview.All()
	detailedReviews := detailedReview.All()

	allReviews := make([]interface{}, len(generalReviews)+len(detailedReviews))

	for x, review := range generalReviews {
		allReviews[x] = review
	}

	for x, review := range detailedReviews {
		allReviews[x+len(generalReviews)] = review
	}

	ctx.JSON(allReviews)
}

// All general reviews
func GeneralReviews(ctx iris.Context) {
	ctx.JSON(generalReview.All())
}

// All detailed reviews
func DetailedReviews(ctx iris.Context) {
	ctx.JSON(detailedReview.All())
}

// All reviews by institution
func ReviewsByInstitution(ctx iris.Context) {
	institution_id, err := ctx.Params().GetInt("id")

	if err != nil {
		fmt.Println(err)
	}

	allReviews := institution.Find(institution_id).Reviews()
	ctx.JSON(allReviews)
}
