package models

import (
    "github.com/jinzhu/gorm"
)

type DetailedReview struct {
    gorm.Model
    Course      string  `gorm:"size:20"`
    Suggestion  string  `gorm:"size:100"`
    Rate        int     `gorm:""`
    // Student relationship - belongs to
    StudentID int `gorm:"ForeignKey:StudentID"`
    // Institution relationship - belongs to
    InstitutionID int `gorm:"ForeignKey:InstitutionID"`
}
