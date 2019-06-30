package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"review-gorm/controllers"
)

func main() {
	e := echo.New()
	e.GET("/:website/review", controllers.GetReviews)
	e.GET("/:website/review/product", controllers.GetProductReviews)
	e.Logger.Fatal(e.Start(":8080"))
}
