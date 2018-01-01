package models

import (
	"github.com/jinzhu/gorm"
)

type GeneralReview struct {
	gorm.Model
	Course      string  `gorm:"size:20"`
	Suggestion  string  `gorm:"size:100"`
	Rate        int     `gorm:""`
	Pros        string  `gorm:"size:60"`
	Cons        string  `gorm:"size:60"`
	Description string  `gorm:"size:160"`
	Payment     float32 `gorm:""`
	StartYear   int     `gorm:""`
	// Student relationship - belongs to
	StudentID int `gorm:"ForeignKey:StudentID"`
	// Institution relationship - belongs to
	InstitutionID int `gorm:"ForeignKey:InstitutionID"`
}
