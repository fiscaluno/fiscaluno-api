package detailedReview

import (
	models "github.com/fiscaluno/fiscaluno-api/models"
	"github.com/mitchellh/mapstructure"
)

type DetailedReview models.DetailedReview

// database instance
var db = models.DB

func (detailed *DetailedReview) Student() (student models.Student) {
	db.Model(&detailed).Related(&student)
	return
}

func (detailed *DetailedReview) Institution() (institution models.Institution) {
	db.Model(&detailed).Related(&institution)
	return
}

// Creates new model of detailed review
// Return detailed review
func Create(attributes map[string]interface{}) *DetailedReview {
	var review DetailedReview
	mapstructure.Decode(attributes, &review)
	return &review
}

// Save new detailed review at database
func (review *DetailedReview) Save() {
	db.Create(&review)
}

// Find detailed review by id
func Find(id int) (review *DetailedReview) {
	review = &DetailedReview{}
	db.First(review, id)
	return
}

// All detailed reviews
func All() (reviews []DetailedReview) {
	db.Find(&reviews)
	return
}
