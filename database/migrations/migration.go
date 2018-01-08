package migrations

import (
	"fmt"
	"github.com/fiscaluno/fiscaluno-api/database"
	log "github.com/fiscaluno/fiscaluno-api/fiscalog"
	models "github.com/fiscaluno/fiscaluno-api/models"
	"reflect"
	"regexp"
)

// Database Instance
var db = database.GetInstance()

func init() {
	log.Info("Migrating database... Wait a moment please")
}

func Migrate() {
	var DetailedReview = models.DetailedReview{}
	var GeneralReview = models.GeneralReview{}
	var Institution = models.Institution{}
	var Student = models.Student{}

	Create(&DetailedReview)
	Create(&GeneralReview)
	Create(&Institution)
	Create(&Student)
}

// Create new Schema to database
func Create(Schema interface{}) {
	pattern := regexp.MustCompile(`\.(\w+)$`)
	schemaString := reflect.TypeOf(Schema).String()
	tableName := pattern.FindStringSubmatch(schemaString)[1]

	if !db.HasTable(Schema) {
		log.Info(fmt.Sprintf("Creating %s table...", tableName))
		defer log.Info(fmt.Sprintf("%s table created", tableName))
		db.AutoMigrate(Schema)
	} else {
		log.Info(fmt.Sprintf("%s table exists...", tableName))
	}
}
