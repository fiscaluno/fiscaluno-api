package institution

import (
	models "github.com/fiscaluno/fiscaluno-api/models"
	"github.com/mitchellh/mapstructure"
)

type Institution models.Institution

// database instance
var db = models.DB

// Institution has many students
func (institution *Institution) Students() (students []models.Student) {
	db.Model(&institution).Related(&students)
	return
}

func (institution *Institution) Reviews() []interface{} {
	general := institution.GeneralReviews()
	detailed := institution.DetailedReviews()
	reviews := make([]interface{}, len(general)+len(detailed))

	for x, review := range general {
		reviews[x] = review
	}

	for x, review := range detailed {
		reviews[x+len(general)] = review
	}

	return reviews
}

// Institution has many general reviews
func (institution *Institution) GeneralReviews() (reviews []models.GeneralReview) {
	db.Model(&institution).Related(&reviews)
	return
}

// Institution has many detailed reviews
func (institution *Institution) DetailedReviews() (reviews []models.DetailedReview) {
	db.Model(&institution).Related(&reviews)
	return
}

// Creates new model of institution
// Return institution
func Create(attributes map[string]interface{}) *Institution {
	var institution Institution
	mapstructure.Decode(attributes, &institution)
	return &institution
}

// Save new institution at database
func (institution *Institution) Save() {
	db.Create(&institution)
}

// Find institution by id
func Find(id int) (institution *Institution) {
	institution = &Institution{}
	db.First(institution, id)
	return
}

// All institutions
func All() (institutions []Institution) {
	db.Find(&institutions)
	return
}
