package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetSingleProduct(t *testing.T) {
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
			if err := GetSingleProduct(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("GetSingleProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateProduct(t *testing.T) {
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
			if err := CreateProduct(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
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
			if err := DeleteProduct(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
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
			if err := GetAllProducts(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("GetAllProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBody(t *testing.T) {
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
			if err := GetBody(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("GetBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetHome(t *testing.T) {
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
			if err := GetHome(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("GetHome() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
