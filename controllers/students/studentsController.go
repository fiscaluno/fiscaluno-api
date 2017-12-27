package students

import (
    "github.com/kataras/iris"
    _"github.com/fiscaluno/fiscaluno-api/database/migrations"
    models "github.com/fiscaluno/fiscaluno-api/models"
)

func InsertStudent(ctx iris.Context) {
    var attributes = make(map[string]interface{})
    values := ctx.FormValues()
    
    for key, value := range values{
        attributes[key] = value[0]
    }

    models.CreateStudent(attributes).Save()
}
