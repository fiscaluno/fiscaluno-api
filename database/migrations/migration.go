package migrations

import (
    "github.com/fiscaluno/fiscaluno-api/database"
)

// Database Instance
var db = database.GetInstance()

// Create new Schema to database
func Create(Schema interface {}) {
    db.AutoMigrate(Schema)
}
