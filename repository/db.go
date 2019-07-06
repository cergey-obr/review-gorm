package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDatabase(dataSource string) {
	var err error

	db, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	db.SingularTable(true)
}
