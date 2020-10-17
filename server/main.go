package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/twoshark/fixture_configuration_server/server/devices"
	"github.com/twoshark/fixture_configuration_server/server/routes"
)

func main() {
	devices := devices.NewDevices("Demo")
	router := mux.NewRouter().StrictSlash(true)
	routes.AddRoutes(router, &devices)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
