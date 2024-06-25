package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
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

func TestGetHome(t *testing.T) {
	tests := []struct {
		description    string
		route          string
		statusCode     int
	}{
		{
			description: "get HTTP status 200",
			route: "/",
			statusCode: 200,
		},
		// Second test case
    {
      description:  "get HTTP status 404, when route is not exists",
      route:        "/not-found",
      statusCode:   404,
    },
	}

	app := fiber.New()
	app.Get("/", GetHome)
	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equal(t, tt.statusCode, resp.StatusCode, tt.description)
	}
}
