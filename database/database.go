package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func GetInstance() (*gorm.DB) {
    db, err := gorm.Open("mysql", "user:password@/database_name?charset=utf8&parseTime=True&loc=Local")
    
    if err != nil {
        fmt.Print("error at database connection")
    }

    return db
}