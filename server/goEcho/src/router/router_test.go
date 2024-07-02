package router

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRegisterRoutes(t *testing.T) {
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
			if got := RegisterRoutes(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
