package repository

import "review-gorm/models"

func GetReviews(filters *models.Filters) []models.Review {
	tx := db.
		Preload("Author").
		Preload("Photos").
		Where("website = ?", filters.Website).
		Limit(filters.Limit).
		Offset(filters.Offset).
		Order("id")

	if len(filters.ProductIds) > 0 {
		tx.Where("product_id IN (?)", filters.ProductIds)
	}

	var reviews []models.Review
	tx.Find(&reviews)

	return reviews
}

func GetProductReviews(filters *models.Filters) models.ProductReviews {
	var totals models.Totals
	db.
		Model(models.Review{}).
		Select("AVG(overall) AS value, COUNT(id) AS count, ROUND(AVG(overall)) AS intval").
		Where("product_id IN (?)", filters.ProductIds).
		Where("website = ?", filters.Website).
		Limit(filters.Limit).
		Offset(filters.Offset).
		Scan(&totals)

	return models.ProductReviews{
		Items:  GetReviews(filters),
		Totals: totals,
	}
}
