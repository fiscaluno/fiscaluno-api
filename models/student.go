package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"github.com/fiscaluno/fiscaluno-api/database"
)

var db = database.GetInstance()

type Student struct {
	gorm.Model
	Name        string `gorm:"size:50"` 
	Phone       string    `gorm:"size:10"`
	Institution string    `gorm:"size:50"`
	Course      string    `gorm:"size:50"`
	Email       string    `gorm:"size:50"`
}

func CreateStudent(studentAttributes map[string] interface {}) (*Student) {
	var student Student
	mapstructure.Decode(studentAttributes, &student)
	
	return &student
}

func (student *Student) Save() {
	db.Create(&student)
	defer db.Close()
}
