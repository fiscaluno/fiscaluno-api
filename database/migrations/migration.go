package migrations

import (
    "fmt"
    "regexp"
    "reflect"
    "github.com/fiscaluno/fiscaluno-api/database"
    log "github.com/fiscaluno/fiscaluno-api/fiscalog"
)

// Database Instance
var db = database.GetInstance()

func init() {
    log.Info.Println("Migrating database... Wait a moment please")
}

func Migrate() {
    log.Info.Println("Database migrated")
}

// Create new Schema to database
func Create(Schema interface {}) {
    pattern := regexp.MustCompile(`\.(\w+)$`)
    schemaString := reflect.TypeOf(Schema).String()
    tableName := pattern.FindStringSubmatch(schemaString)[1]

    if !db.HasTable(Schema) {
        log.Info.Println(fmt.Sprintf("Creating %s table...", tableName))
        defer log.Info.Println(fmt.Sprintf("%s table created", tableName))
        db.AutoMigrate(Schema)
    } else {
        log.Info.Println(fmt.Sprintf("%s table exists...", tableName))
    }
}
