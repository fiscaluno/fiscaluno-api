package migrations

import (
    "github.com/fiscaluno/fiscaluno-api/models"
)

var Student = models.Student{}

func init() {
    Create(&Student)
}
