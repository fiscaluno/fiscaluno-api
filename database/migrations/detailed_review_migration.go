package migrations

import (
	model "github.com/fiscaluno/fiscaluno-api/models"
)

var DetailedReview = model.DetailedReview{}

func init() {
	Create(&DetailedReview)
}
