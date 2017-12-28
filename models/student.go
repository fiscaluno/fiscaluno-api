package models

import (
	"github.com/jinzhu/gorm"
)

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
	InstitutionID int `gorm:"not null" json:"InstitutionId"`

	// Course Relationship - Belongs To
	// CourseID      int    `gorm:"not null" json:"CourseId"`
}
