package middlewares

import (
	// "reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLogger(t *testing.T) {
	type args struct {
		context *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Logger(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("Logger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNext(t *testing.T) {
	type args struct {
		context *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Next(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
