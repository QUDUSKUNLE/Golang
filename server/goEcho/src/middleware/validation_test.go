package middleware

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func TestCustomValidator_Validate(t *testing.T) {
	type fields struct {
		validator *validator.Validate
	}
	type args struct {
		inter interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := &CustomValidator{
				validator: tt.fields.validator,
			}
			if err := cv.Validate(tt.args.inter); (err != nil) != tt.wantErr {
				t.Errorf("CustomValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegisterValidation(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
		want *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterValidation(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
