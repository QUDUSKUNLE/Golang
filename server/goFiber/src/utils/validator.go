package utils

import (
	"github.com/go-playground/validator/v10"
  "github.com/google/uuid"
)

func Validator() *validator.Validate {
	// Validate a model
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
				return true
		}
		return false
	})
	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
			fields[err.Field()] = err.Error()
	}
	return fields
}
