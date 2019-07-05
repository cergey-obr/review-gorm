package main

import (
	"github.com/labstack/echo"
	"review-gorm/controllers"
	"review-gorm/repository"
)

func main() {
	repository.InitDatabase()

	e := echo.New()
	e.GET("/:website/review", controllers.GetReviews)
	e.GET("/:website/review/product", controllers.GetProductReviews)
	e.Logger.Fatal(e.Start(":8080"))
}
