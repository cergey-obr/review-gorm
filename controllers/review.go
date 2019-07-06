package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"review-gorm/models"
	"review-gorm/repository"
)

func GetReviews(c echo.Context) error {
	filters := new(models.Filters)
	err := c.Bind(filters)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, repository.GetReviews(filters))
}

func GetProductReviews(c echo.Context) error {
	filters := new(models.Filters)
	err := c.Bind(filters)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, repository.GetProductReviews(filters))
}
