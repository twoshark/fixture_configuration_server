package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/twoshark/fixture_configuration_server/server/devices"
	"github.com/twoshark/fixture_configuration_server/server/fixtures"
)

//AddFixtureHandlers ...
func AddFixtureHandlers(e *echo.Echo, d *devices.Devices) {
	e.GET("/devices/:device/fixture/:name", GetFixture(d))
	e.POST("/devices/:device/fixture/:name", CreateFixture(d))
	e.PUT("/devices/:device/fixture/:name", UpdateFixture(d))
	e.DELETE("/devices/:device/fixture/:name", DeleteFixture(d))
}

//GetFixture ...
func GetFixture(d *devices.Devices) func(c echo.Context) error {
	fmt.Println("Endpoint Hit: Get Fixture")
	return func(c echo.Context) error {
		name := c.Param("name")
		deviceName := c.Param("device")
		_, device := d.Find(deviceName)
		_, fixture := device.BasicFixtures.Find(name)
		return c.JSON(http.StatusOK, fixture)
	}
}

//CreateFixture ...
func CreateFixture(d *devices.Devices) func(c echo.Context) error {
	fmt.Println("Endpoint Hit: Create Fixture")
	return func(c echo.Context) error {
		var newFixture fixtures.BasicFixture
		if err := c.Bind(newFixture); err != nil {
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}
		deviceName := c.Param("device")
		_, device := d.Find(deviceName)
		index, fixture := device.BasicFixtures.Find(newFixture.Name)
		if index != -1 {
			return c.JSON(http.StatusBadRequest, "ERROR: This Fixture Already Exists")
		}
		device.BasicFixtures.Add(fixture)
		return c.JSON(http.StatusOK, fixture)
	}
}

//UpdateFixture ...
func UpdateFixture(d *devices.Devices) func(c echo.Context) error {
	fmt.Println("Endpoint Hit: Update Fixture")
	return func(c echo.Context) error {
		var updateFixture fixtures.BasicFixture
		if err := c.Bind(updateFixture); err != nil {
			return c.JSON(http.StatusBadRequest, updateFixture)
		}

		deviceName := c.Param("device")
		_, device := d.Find(deviceName)
		index, fixture := device.BasicFixtures.Find(updateFixture.Name)
		if index == -1 {
			return c.JSON(http.StatusBadRequest, "ERROR: The Fixture Was Not Found. Try Creating")
		}
		device.BasicFixtures.Update(index, fixture)
		return c.JSON(http.StatusOK, device.BasicFixtures[index])
	}
}

//DeleteFixture ...
func DeleteFixture(d *devices.Devices) func(c echo.Context) error {
	fmt.Println("Endpoint Hit: Delete Fixture")
	return func(c echo.Context) error {

		var deleteFixture fixtures.BasicFixture
		if err := c.Bind(deleteFixture); err != nil {
			return c.JSON(http.StatusBadRequest, "Bad Requesr")
		}

		deviceName := c.Param("device")
		_, device := d.Find(deviceName)
		index, _ := device.BasicFixtures.Find(deleteFixture.Name)
		if index != -1 {
			device.BasicFixtures.Delete(index, true)
			return c.JSON(http.StatusOK, "Deleted.")
		}
		return c.JSON(http.StatusNotFound, "Not Found.")
	}
}
