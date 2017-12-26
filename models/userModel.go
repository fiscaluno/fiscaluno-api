package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Age  int
	Name string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Num  int    `gorm:"AUTO_INCREMENT"`
}

func conectDB() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "jc:totvs@123@/fiscaluno?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func MigrateUser() {
	db := conectDB()
	db.AutoMigrate(&User{})
	defer db.Close()
}

func CreateUser() {
	MigrateUser()
	user := User{Name: "Jinzhu", Age: 18}
	db := conectDB()
	db.Create(&user)
	defer db.Close()
}
