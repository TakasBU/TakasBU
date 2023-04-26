package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET("/hello", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "hello world") //TODO DEGİSTİR
	})

	e.Static("/static", "static")
}
