package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/fiscaluno/fiscaluno-api/config"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func GetInstance() (*gorm.DB) {
    if db == nil { 
        var err interface {}
        
        user := config.GetConfig("db_user")
        password := config.GetConfig("db_password")
        database := config.GetConfig("db_name")

        connect := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, database)

        db, err = gorm.Open("mysql", connect)
        
        if err != nil {
            panic("error at database connection")
        }
    }

    return db
}