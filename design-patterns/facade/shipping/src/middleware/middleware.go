package middleware

import (
	"net/http"
	// "github.com/labstack/echo/v4/middleware"
  "gopkg.in/go-playground/validator.v9"
	"github.com/labstack/echo/v4"
)

// Custom validator
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(inter interface{}) error {
	if err := cv.validator.Struct(inter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ValidationAdaptor(e *echo.Echo) *echo.Echo {
	e.Validator = &CustomValidator{validator: validator.New()}
	return e
}
