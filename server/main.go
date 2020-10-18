package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/twoshark/fixture_configuration_server/server/devices"
	"github.com/twoshark/fixture_configuration_server/server/routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	devices := devices.NewDevices("Demo")
	routes.AddRoutes(e, &devices)
	log.Println("Start Server")
	log.Fatal(e.Start("0.0.0.0:10000"))
}
