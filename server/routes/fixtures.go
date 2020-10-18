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

	fixturesGroup := e.Group("api/devices/:device/fixtures")
	fixturesGroup.GET("/", getFixtures(d))

	fixtureGroup := e.Group("api/devices/:device/fixtures/:name")
	fixtureGroup.GET("/", GetFixture(d))
	fixtureGroup.POST("/", CreateFixture(d))
	fixtureGroup.PUT("/", UpdateFixture(d))
	fixtureGroup.DELETE("/", DeleteFixture(d))
}

//GetFixture ...
func GetFixture(d *devices.Devices) func(c echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Get Fixture")
		name := c.Param("name")
		deviceName := c.Param("device")
		_, device := d.Find(deviceName)
		_, fixture := device.BasicFixtures.Find(name)
		return c.JSON(http.StatusOK, fixture)
	}
}

//GetFixtures ...
func getFixtures(d *devices.Devices) func(c echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Get Fixture")
		deviceName := c.Param("device")
		_, device := d.Find(deviceName)
		return c.JSON(http.StatusOK, device.BasicFixtures)
	}
}

//CreateFixture ...
func CreateFixture(d *devices.Devices) func(c echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Create Fixture")
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
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Update Fixture")
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
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Delete Fixture")
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
