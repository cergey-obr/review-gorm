package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"os"
	"review-gorm/controllers"
	"review-gorm/repository"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Error("Error loading .env file")
	}

	repository.InitDatabase(os.Getenv("DATABASE"))

	e.GET("/:website/review", controllers.GetReviews)
	e.GET("/:website/review/product", controllers.GetProductReviews)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
