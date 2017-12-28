package student

import (
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"github.com/fiscaluno/fiscaluno-api/models"
	"github.com/fiscaluno/fiscaluno-api/models/institution"
)

var db = models.DB

type Student struct {
	gorm.Model
	// TODO: Relationships ( Institution and Course )
	Name        string `gorm:"size:50" json:"Name"`
	Email       string `gorm:"size:40;unique" json:"Email"`
	Gender      string `gorm:"size:10" json:"Gender"`
	Nacionality string `gorm:"size:15" json:"Nacionality"`
	CivilStatus string `gorm:"size:10" json:"CivilStatus"`
	Address     string `gorm:"size:50" json:"CivilStatus"`
	City        string `gorm:"size:15" json:"City"`
	State       string `gorm:"size:15" json:"State"`
	Phone       string `gorm:"size:10" json:"Phone"`

	// Insitution Relationship - Belongs To
	institution   int `gorm:"ForeignKey:InstitutionID"`
	InstitutionID int `gorm:"not null;ForeignKey:" json:"InstitutionId"`

	// Course Relationship - Belongs To
	// CourseID      int    `gorm:"not null" json:"CourseId"`
}

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

// Retrieves student model
// Return student
func Find(id int) (student *Student) {
	student = &Student{}
	db.First(student, id)
	return
}

// Save new student at database
func (student *Student) Save() {
	db.Create(&student)
}
