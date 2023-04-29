package main

import (
	"log"

	"github.com/TakasBU/TakasBU/initializers"
	"github.com/TakasBU/TakasBU/middlewares"
	"github.com/TakasBU/TakasBU/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	e := echo.New()

	middlewares.Middleware(e)

	routes.Route(e)

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
