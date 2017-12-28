package migrations

import (
	model "github.com/fiscaluno/fiscaluno-api/models"
)

var Student = model.Student{}

func init() {
	Create(&Student)
}
