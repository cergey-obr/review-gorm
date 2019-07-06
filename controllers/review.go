package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"review-gorm/middleware"
	"review-gorm/repository"
)

func GetReviews(c echo.Context) error {
	controllerContext := c.(*middleware.ControllerContext)
	return c.JSON(http.StatusOK, repository.GetReviews(controllerContext.Filters))
}

func GetProductReviews(c echo.Context) error {
	controllerContext := c.(*middleware.ControllerContext)
	return c.JSON(http.StatusOK, repository.GetProductReviews(controllerContext.Filters))
}
