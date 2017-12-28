package migrations

import (
    model "github.com/fiscaluno/fiscaluno-api/models"
)

var Institution = model.Institution{}

func init() {
    Create(&Institution)
}
