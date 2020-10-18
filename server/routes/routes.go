package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/twoshark/fixture_configuration_server/server/devices"
)

//AddRoutes ...
func AddRoutes(e *echo.Echo, d *devices.Devices) {
	e.GET("/", home)
	AddDevicesRoutes(e, d)
	addFixtureHandlers(e, d)

}

func home(c echo.Context) error {
	fmt.Println("Endpoint Hit: homePage")
	return c.JSON(http.StatusOK, "Welcome to The Installation Config Server")
}
