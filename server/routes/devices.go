package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/twoshark/fixture_configuration_server/server/devices"
)

//AddDevicesRoutes ...
func AddDevicesRoutes(router *mux.Router, d *devices.Devices) {
	router.HandleFunc("/devices", GetDevices(d))
	router.HandleFunc("/devices/{name}", GetDevice(d))
	//router.HandleFunc("/devices", GetInstallations(&i)).Methods("POST")
	//router.HandleFunc("/devices", GetInstallations(&i)).Methods("PUT")
	//router.HandleFunc("/devices", GetInstallations(&i)).Methods("DELETE")
}

//GetDevices ...
func GetDevices(d *devices.Devices) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetInstallations")
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(d)
	}
}

//GetDevice ...
func GetDevice(d *devices.Devices) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetDevice")
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		index, device := d.Find(name)
		if index >= 0 {
			json.NewEncoder(w).Encode(device)
		}
	}
}
