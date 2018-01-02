package students

import (
    "strconv"
    "reflect"
	student "github.com/fiscaluno/fiscaluno-api/models/student"
	"github.com/kataras/iris"
)

// Creates new student - POST Method
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

	student.Create(attributes).Save()
}

// Find student at database - GET Method
func FindById(ctx iris.Context) {
    id, err := ctx.Params().GetInt("id")
    
    if err != nil {
        ctx.JSON(err)
    }
    
    // Checks if student exists at database
    if reflect.DeepEqual(student.Student{}, *student.Find(id)) {
        ctx.StatusCode(404)
    } else {
        ctx.JSON(student.Find(id))    
    }
}