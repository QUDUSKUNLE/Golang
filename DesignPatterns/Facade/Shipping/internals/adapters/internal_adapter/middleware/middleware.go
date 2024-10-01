package middleware

import (
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Custom validator
type CustomValidator struct {
	validator *validator.Validate
}

// Custom validator
func (c *CustomValidator) Validate(inter interface{}) error {
	if err := c.validator.Struct(inter); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors))
		}
		var errorMessage []map[string]string
		for _, er := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, map[string]string{"field": er.Field(), "message": fmt.Sprintf("%s is an invalid input for field: %s", er.Value(), er.Field())})
		}
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}
	return nil
}

// ValidationAdaptor
func ValidationAdaptor(e *echo.Echo) *echo.Echo {
	e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
	return e
}
