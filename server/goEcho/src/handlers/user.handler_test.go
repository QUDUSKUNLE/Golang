package handlers

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHome(t *testing.T) {
	type args struct {
		con echo.Context
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
			if err := Home(tt.args.con); (err != nil) != tt.wantErr {
				t.Errorf("Home() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPong(t *testing.T) {
	type args struct {
		con echo.Context
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
			if err := Pong(tt.args.con); (err != nil) != tt.wantErr {
				t.Errorf("Pong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		con echo.Context
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
			if err := Register(tt.args.con); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
