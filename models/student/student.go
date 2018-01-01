package student

import (
	models "github.com/fiscaluno/fiscaluno-api/models"
	"github.com/fiscaluno/fiscaluno-api/models/institution"
	"github.com/mitchellh/mapstructure"
)

type Student models.Student

// database instance
var db = models.DB

// Student belongs to one institution
// Return institution
func (student *Student) Institution() (studentInstitution *institution.Institution) {
	studentInstitution = &institution.Institution{}
	db.Model(&student).Related(studentInstitution)
	return
}

// Creates new model of student
// Return student
func Create(attributes map[string]interface{}) *Student {
	var student Student
	mapstructure.Decode(attributes, &student)
	return &student
}

// Save new student at database
func (student *Student) Save() (err error) {
	if rows := db.Create(&student).RowsAffected; rows != 1 {
		return err
	}
	return
}

// Find student by id
// Return student
func Find(id int) (student *Student) {
	student = &Student{}
	db.First(student, id)
	return
}
