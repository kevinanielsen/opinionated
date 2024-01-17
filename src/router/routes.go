package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
