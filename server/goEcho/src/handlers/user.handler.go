package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/QUDUSKUNLE/goecho/src/models"
)

func Home(con echo.Context) error {
	return con.JSON(http.StatusOK, map[string]string{"message": "Hello World!"})
}

func Pong(con echo.Context) error {
	return con.JSON(http.StatusOK, map[string]string{"ping": "pong"})
}

func Register(con echo.Context) error {
	user := models.User{}
	if err := con.Bind(&user); err != nil {
		return con.JSON(http.StatusBadRequest, err)
	}

	// Validate user input
	if err := con.Validate(&user); err != nil {
		return con.JSON(http.StatusBadRequest, err)
	}
	// Process valid user data
	return con.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}

