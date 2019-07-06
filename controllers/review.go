package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"review-gorm/repository"
	"strconv"
)

func GetReviews(c echo.Context) error {
	productIds := make([]string, 0)
	website, _ := strconv.Atoi(c.Param("website"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	return c.JSON(http.StatusOK, repository.GetReviews(website, productIds, limit, offset))
}

func GetProductReviews(c echo.Context) error {
	params := c.QueryParams()

	website, _ := strconv.Atoi(c.Param("website"))
	productIds := params["productId[]"]
	limit, _ := strconv.Atoi(params.Get("limit"))
	offset, _ := strconv.Atoi(params.Get("offset"))

	return c.JSON(http.StatusOK, repository.GetProductReviews(website, productIds, limit, offset))
}
