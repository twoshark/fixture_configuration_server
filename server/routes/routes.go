package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/twoshark/fixture_configuration_server/server/devices"
)

//AddRoutes ...
func AddRoutes(router *mux.Router, d *devices.Devices) {
	router.HandleFunc("/", home)
	router.HandleFunc("/refresh", refresh)
	AddDevicesRoutes(router, d)
	for _, device := range d.Devices {
		addFixtureHandlers(router, &device)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to The Installation Config Server")
	fmt.Println("Endpoint Hit: homePage")
}

func refresh(w http.ResponseWriter, r *http.Request) {

}
