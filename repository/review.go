package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"review-gorm/models"
)

func GetAllReviews(website int, limit int, offset int) []models.Review {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/review")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//db.LogMode(true)
	db.SingularTable(true)

	var reviews []models.Review
	db.
		Preload("Author").
		Preload("Photos").
		Where("website = ?", website).
		Limit(limit).
		Offset(offset).
		Find(&reviews)

	return reviews
}

func GetProductReviews(website int, productIds []string, limit int, offset int) []models.Review {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/review")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.LogMode(true)
	db.SingularTable(true)

	var reviews []models.Review
	db.
		Preload("Author").
		Preload("Photos").
		Where("website = ?", website).
		Where("product_id IN (?)", productIds).
		Limit(limit).
		Offset(offset).
		Find(&reviews)

	return reviews
}
