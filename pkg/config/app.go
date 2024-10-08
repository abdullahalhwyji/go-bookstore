package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", `root:@tcp(localhost:3306)/books?charset=utf8&parseTime=True&loc=Local`)

	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("The application is running")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
