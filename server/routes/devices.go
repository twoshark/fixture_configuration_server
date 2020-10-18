package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/twoshark/fixture_configuration_server/server/devices"
)

//AddDevicesRoutes ...
func AddDevicesRoutes(router *mux.Router, d *devices.Devices) {
	router.HandleFunc("/devices", GetDevices(d))
	router.HandleFunc("/devices/{name}", GetDevice(d))
	//router.HandleFunc("/devices/{name}", CreateDevice(&i)).Methods("POST")
	//router.HandleFunc("/devices/{name}", UpdateDevice(&i)).Methods("PUT")
	//router.HandleFunc("/devices/{name}", DeleteDevice(&i)).Methods("DELETE")
}

//GetDevices ...
func GetDevices(d *devices.Devices) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get Devices")
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(d)
	}
}

//GetDevice ...
func GetDevice(d *devices.Devices) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get Device")
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		index, device := d.Find(name)
		if index >= 0 {
			json.NewEncoder(w).Encode(device)
		}
	}
}

//CreateDevice ...
func CreateDevice(d *devices.Devices) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Create Device")
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var newDevice devices.Device
		json.Unmarshal(reqBody, &newDevice)

		index, _ := d.Find(newDevice.Name)
		if index != -1 {
			fmt.Fprintf(w, "ERROR: This Device Already Exists")
			return
		}
		d.Add(newDevice)
		json.NewEncoder(w).Encode(newDevice)
	}
}

//UpdateDevice ...
func UpdateDevice(d *devices.Devices) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Update Device")
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var update devices.Device
		json.Unmarshal(reqBody, &update)

		index, device := d.Find(update.Name)
		if index == -1 {
			fmt.Fprintf(w, "ERROR: The Device Was Not Found. Try Creating")
		}
		d.Update(index, device)
		json.NewEncoder(w).Encode(d.Devices[index])
	}
}

//DeleteDevice ...
func DeleteDevice(d *devices.Devices) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete Device")
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var delete devices.Device
		json.Unmarshal(reqBody, &delete)

		index, _ := d.Find(delete.Name)
		if index != -1 {
			d.Delete(index, true)
			fmt.Fprintf(w, "Deleted.")
		}
		fmt.Fprintf(w, "Not Found.")
	}
}
