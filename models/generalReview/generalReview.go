package generalReview

import (
	models "github.com/fiscaluno/fiscaluno-api/models"
	"github.com/mitchellh/mapstructure"
)

type GeneralReview models.GeneralReview

// database instance
var db = models.DB

func (general *GeneralReview) Student() (student models.Student) {
	db.Model(&general).Related(&student)
	return
}

func (general *GeneralReview) Institution() (institution models.Institution) {
	db.Model(&general).Related(&institution)
	return
}

// Creates new model of general review
// Return general review
func Create(attributes map[string]interface{}) *GeneralReview {
	var review GeneralReview
	mapstructure.Decode(attributes, &review)
	return &review
}

// Save new general review at database
func (review *GeneralReview) Save() {
	db.Create(&review)
}

// Find general review by id
func Find(id int) (review *GeneralReview) {
	review = &GeneralReview{}
	db.First(review, id)
	return
}

// All general reviews
func All() (reviews []GeneralReview) {
	db.Find(&reviews)
	return
}