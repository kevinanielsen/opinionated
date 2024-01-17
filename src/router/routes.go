package router

import (
	"log"
	"net/http"

	"github.com/kevinanielsen/opinionated/ui"
	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()
	e.StaticFS("/", ui.DistDirFS)

	e.GET("/api", func(c echo.Context) error {
		log.Println("GET /api")
		return c.JSON(http.StatusOK, "Pong!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
