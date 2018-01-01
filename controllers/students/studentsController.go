package students

import (
    "strconv"
	students "github.com/fiscaluno/fiscaluno-api/models/student"
	"github.com/kataras/iris"
)

func InsertStudent(ctx iris.Context) {
	var attributes = make(map[string]interface{})
	values := ctx.FormValues()

	for key, value := range values {
        intVal, err := strconv.Atoi(value[0])
        if err != nil {
          attributes[key] = value[0]
        }else {
		  attributes[key] = intVal
        }
	}

	students.Create(attributes).Save()
}
