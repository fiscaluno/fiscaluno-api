package migrations

import (
    "fmt"
    "regexp"
    "reflect"
    "github.com/fiscaluno/fiscaluno-api/database"
)

// Database Instance
var db = database.GetInstance()

func init() {
    fmt.Println("[AutoMigration: Migrating database... Wait a moment please]")
}

func Migrate() {
    fmt.Println("[AutoMigration: Database migrated]")
}

// Create new Schema to database
func Create(Schema interface {}) {
    pattern := regexp.MustCompile(`\.(\w+)$`)
    schemaString := reflect.TypeOf(Schema).String()
    tableName := pattern.FindStringSubmatch(schemaString)[1]

    if !db.HasTable(Schema) {
        fmt.Println(fmt.Sprintf("[AutoMigration: Creating %s table...]", tableName))
        defer fmt.Println("[AutoMigration: Students table created]")
        db.AutoMigrate(Schema)
    } else {
        fmt.Println(fmt.Sprintf("[AutoMigration: %s table exists...]", tableName))
    }
}
