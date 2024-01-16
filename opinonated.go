package main

import (
	"net/http"

	"github.com/kevinanielsen/opinionated/src/database"
	"github.com/kevinanielsen/opinionated/src/initializers"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
