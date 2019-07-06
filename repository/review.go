package repository

import "review-gorm/models"

func GetReviews(website int, productIds []string, limit int, offset int) []models.Review {
	tx := db.
		Preload("Author").
		Preload("Photos").
		Where("website = ?", website).
		Limit(limit).
		Offset(offset).
		Order("id")

	if len(productIds) > 0 {
		tx.Where("product_id IN (?)", productIds)
	}

	var reviews []models.Review
	tx.Find(&reviews)

	return reviews
}

func GetProductReviews(website int, productIds []string, limit int, offset int) models.ProductReviews {
	var totals models.Totals
	db.
		Model(models.Review{}).
		Select("AVG(overall) AS value, COUNT(id) AS count, ROUND(AVG(overall)) AS intval").
		Where("product_id IN (?)", productIds).
		Where("website = ?", website).
		Limit(limit).
		Offset(offset).
		Scan(&totals)

	return models.ProductReviews{
		Items:  GetReviews(website, productIds, limit, offset),
		Totals: totals,
	}
}
