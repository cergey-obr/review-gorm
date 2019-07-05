package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDatabase() {
	var err error

	db, err = gorm.Open("mysql", "go:123456@tcp(localhost:3306)/review")
	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	db.SingularTable(true)
}
