package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// this whole config package will ultimately return a variable called db
// which will help other files to interact with the database

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "rana:0308@%@/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
