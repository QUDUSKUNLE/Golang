package middleware

import (
	"net/http"
  "gopkg.in/go-playground/validator.v9"
	"github.com/labstack/echo/v4"
)

// Custom validator
type CustomValidator struct {
	validator *validator.Validate
}

func (custom *CustomValidator) Validate(inter interface{}) error {
	if err := custom.validator.Struct(inter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ValidationAdaptor(e *echo.Echo) *echo.Echo {
	e.Validator = &CustomValidator{validator: validator.New()}
	return e
}
