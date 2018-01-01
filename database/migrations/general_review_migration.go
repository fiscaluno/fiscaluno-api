package migrations

import (
	model "github.com/fiscaluno/fiscaluno-api/models"
)

var GeneralReview = model.GeneralReview{}

func init() {
	Create(&GeneralReview)
}
