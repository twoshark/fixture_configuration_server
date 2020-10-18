package main

import (
	"log"
	"net/http"
	"sort"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/twoshark/fixture_configuration_server/server/devices"
	"github.com/twoshark/fixture_configuration_server/server/fixtures"
	"github.com/twoshark/fixture_configuration_server/server/routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//TODO: load Devices from somewhere?
	fxs := make(fixtures.BasicFixtures, 0)
	fxs = append(fxs,
		fixtures.BasicFixture{
			Name: "1",
		},
		fixtures.BasicFixture{
			Name: "2",
		})
	devices := devices.NewDevices("Demo",
		devices.Device{
			Name:          "dev1",
			BasicFixtures: fxs,
		},
		devices.Device{
			Name:          "dev2",
			BasicFixtures: fxs,
		},
		devices.Device{
			Name:          "dev3",
			BasicFixtures: fxs,
		},
	)
	// frontend
	e.Static("/static", "../build/static")
	e.File("/*", "../build/index.html")
	// backend
	routes.AddDevicesRoutes(e, &devices)
	routes.AddFixtureHandlers(e, &devices)

	e.GET("/routes", func(c echo.Context) error {
		var rts []echo.Route
		for _, rt := range e.Routes() {
			rts = append(rts, *rt)
		}
		sort.Sort(SortBy(rts))
		return c.JSON(http.StatusOK, rts)
	})

	log.Println("Start Server...")
	log.Fatal(e.Start("0.0.0.0:10000"))
}

//SortBy ...
type SortBy []echo.Route

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].Path < a[j].Path }
