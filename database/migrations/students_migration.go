package migrations

import (
    model "github.com/fiscaluno/fiscaluno-api/models/student"
)

var Student = model.Student{}

func init() {
    Create(&Student)
}
