package main

import (
	"github.com/TakasBU/TakasBU/middlewares"
	"github.com/TakasBU/TakasBU/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	middlewares.Middleware(e)

	// Routes
	routes.Route(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
