package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/twoshark/fixture_configuration_server/server/devices"
)

//AddDevicesRoutes ...
func AddDevicesRoutes(e *echo.Echo, d *devices.Devices) {
	devicesGroup := e.Group("api/devices")
	devicesGroup.GET("/", GetDevices(d))

	deviceGroup := e.Group("api/devices/:name")
	deviceGroup.GET("/", GetDevice(d))
	deviceGroup.POST("/", CreateDevice(d))
	deviceGroup.PUT("/", UpdateDevice(d))
	deviceGroup.DELETE("/", DeleteDevice(d))

}

//GetDevices ...
func GetDevices(d *devices.Devices) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Get Devices")
		return c.JSON(200, d)
	}
}

//GetDevice ...
func GetDevice(d *devices.Devices) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Get Device")
		name := c.Param("name")
		index, device := d.Find(name)
		if index >= 0 {
			return c.JSON(http.StatusOK, device)
		}
		return c.JSON(http.StatusNotFound, "Not Found")
	}
}

//CreateDevice ...
func CreateDevice(d *devices.Devices) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Create Device")

		var newDevice devices.Device
		if err := c.Bind(newDevice); err != nil {
			return err
		}

		index, _ := d.Find(newDevice.Name)
		if index != -1 {
			return c.JSON(http.StatusNotFound, "ERROR: This Device Already Exists")
		}
		d.Add(newDevice)
		return c.JSON(http.StatusOK, newDevice)
	}
}

//UpdateDevice ...
func UpdateDevice(d *devices.Devices) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Update Device")
		var update devices.Device
		if err := c.Bind(update); err != nil {
			return err
		}

		index, device := d.Find(update.Name)
		if index == -1 {
			return c.JSON(http.StatusNotFound, "ERROR: The Device Was Not Found. Try Creating")
		}
		d.Update(index, device)
		return c.JSON(http.StatusOK, d.Devices[index])
	}
}

//DeleteDevice ...
func DeleteDevice(d *devices.Devices) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Endpoint Hit: Delete Device")
		var delete devices.Device
		if err := c.Bind(delete); err != nil {
			return err
		}

		index, _ := d.Find(delete.Name)
		if index != -1 {
			d.Delete(index, true)

			return c.JSON(http.StatusOK, "Deleted.")
		}
		return c.JSON(http.StatusNotFound, "Not Found")
	}
}
