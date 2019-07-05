package repository

import "review-gorm/models"

func GetAllReviews(website int, limit int, offset int) []models.Review {
	if limit == 0 {
		limit = 10
	}

	var reviews []models.Review
	db.
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
	var reviews []models.Review
	db.
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
