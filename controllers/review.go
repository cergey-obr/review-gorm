package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"review-gorm/repository"
	"strconv"
)

func GetReviews(c echo.Context) error {
	website, _ := strconv.Atoi(c.Param("website"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	filter := c.QueryParam("filter")
	fmt.Println(filter)

	return c.JSON(http.StatusOK, repository.GetAllReviews(website, limit, offset))
}
