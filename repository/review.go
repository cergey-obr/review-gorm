package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"review-gorm/models"
)

var DBCon *gorm.DB

func init() {
	DBCon = connect()
}
func connect() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:123456@tcp(localhost:3307)/review")
	return db
}

func GetAllReviews(website int, limit int, offset int) []models.Review {
	if limit == 0 {
		limit = 10
	}

	DBCon.LogMode(true)
	DBCon.SingularTable(true)

	var reviews []models.Review
	DBCon.
		Preload("Author").
		Preload("Photos").
		Where("website = ?", website).
		Limit(limit).
		Offset(offset).
		Order("id").
		Find(&reviews)

	return reviews
}

func GetProductReviews(website int, productIds []string, limit int, offset int) models.ProductReviews {
	DBCon.LogMode(true)
	DBCon.SingularTable(true)

	var reviews []models.Review
	DBCon.
		Preload("Author").
		Preload("Photos").
		Where("product_id IN (?)", productIds).
		Where("website = ?", website).
		Limit(limit).
		Offset(offset).
		Order("id").
		Find(&reviews)

	var productReviews models.ProductReviews
	var totals models.Totals
	DBCon.
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
