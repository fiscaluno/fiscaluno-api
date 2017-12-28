package institution

import (
    "github.com/jinzhu/gorm"
    "github.com/mitchellh/mapstructure"
	"github.com/fiscaluno/fiscaluno-api/models"
)

var db = models.DB

type Institution struct {
	gorm.Model
	Name     string `gorm:"size:50",json:Name`
	Address  string `gorm:"size:50",json:Address`
	Cnpj     string `gorm:"size:14",json:Cnpj`
	Email    string `gorm:"size:40",json:Email`
	Website  string `gorm:"size:30",json:Website`
	Phone    string `gorm:"size:10",json:Phone`
	ImageUri string `gorm:"size:50",json:ImageUri`
}

func Create(attributes map[string]interface{}) *Institution {
	var institution Institution
	mapstructure.Decode(attributes, &institution)
	return &institution
}

func (institution *Institution) Save() {
	db.Create(&institution)
}
