package database

import (
	"fmt"
	"github.com/fiscaluno/fiscaluno-api/config"
	log "github.com/fiscaluno/fiscaluno-api/fiscalog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		var err interface{}

		db, err = gorm.Open("mysql", getConnectionString())

		if err != nil {
			log.Error(err)
			panic("error at database connection")
		}
	}

	return db
}

func getConnectionString() (connect string) {
	user := config.GetConfig("db_user")
	password := config.GetConfig("db_password")
	database := config.GetConfig("db_name")
	connect = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, database)
	return
}
