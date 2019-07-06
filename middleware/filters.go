package middleware

import (
	"github.com/labstack/echo"
	"review-gorm/models"
)

type ControllerContext struct {
	echo.Context
	Filters *models.Filters
}

func AddFilters(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		filters := new(models.Filters)
		err := context.Bind(filters)
		if err != nil {
			return err
		}

		controllerContext := &ControllerContext{context, filters}
		return next(controllerContext)
	}
}
