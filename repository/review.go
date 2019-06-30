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

	db.LogMode(true)
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

func GetProductReviews(website int, productIds []string, limit int, offset int) models.ProductReviews {
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
		Limit(limit).
		Offset(offset).
		Find(&reviews)

	var productReviews models.ProductReviews
	var totals models.Totals
	db.
		Model(models.Review{}).
		Select("AVG(overall) AS value, COUNT(id) AS count, ROUND(AVG(overall)) AS intval").
		Where("product_id IN (?)", productIds).
		Where("website = ?", website).
		Limit(limit).
		Offset(offset).
		Scan(&totals)

	productReviews.Items = reviews
	productReviews.Totals = totals

	return productReviews
}
