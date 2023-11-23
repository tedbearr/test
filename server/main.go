package main

import (
	"net/http"
	"server/config"
	"server/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	config.Env()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	prefix := e.Group("/api/v1/")
	route.AuthRoute(prefix)

	e.GET("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, "not found")
	})

	e.Logger.Fatal(e.Start(":" + "8070"))
}
