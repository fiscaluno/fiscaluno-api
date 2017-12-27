package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

type Student struct {
	gorm.Model
	Name        string `gorm:"size:50",json:Name`
	Phone       string `gorm:"size:10",json:Phone`
	Institution string `gorm:"size:50",json:Institution`
	Course      string `gorm:"size:50",json:Course`
	Email       string `gorm:"size:50",json:Email`
}

func CreateStudent(studentAttributes map[string]interface{}) *Student {
	var student Student
	mapstructure.Decode(studentAttributes, &student)
	return &student
}

func (student *Student) Save() {
	db.Create(&student)
}
