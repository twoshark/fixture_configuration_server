package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/twoshark/fixture_configuration_server/server/devices"
	"github.com/twoshark/fixture_configuration_server/server/fixtures"
)

func addFixtureHandlers(router *mux.Router, d *devices.Device) {
	router.HandleFunc("/devices/"+d.Name+"/fixture/{name}", GetFixture(d))
	router.HandleFunc("/devices/"+d.Name+"/fixture/{name}", CreateFixture(d)).Methods("POST")
	router.HandleFunc("/devices/"+d.Name+"/fixture/{name}", UpdateFixture(d)).Methods("PUT")
	router.HandleFunc("/devices/"+d.Name+"/fixture/{name}", DeleteFixture(d)).Methods("DELETE")
}

//GetFixture ...
func GetFixture(d *devices.Device) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get Fixture")
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		_, fixture := d.BasicFixtures.Find(name)
		json.NewEncoder(w).Encode(fixture)
	}
}

//CreateFixture ...
func CreateFixture(d *devices.Device) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Create Fixture")
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var newFixture fixtures.BasicFixture
		json.Unmarshal(reqBody, &newFixture)

		index, fixture := d.BasicFixtures.Find(newFixture.Name)
		if index != -1 {
			fmt.Fprintf(w, "ERROR: This Fixture Already Exists")
		}
		d.BasicFixtures.Add(fixture)
		json.NewEncoder(w).Encode(fixture)
	}
}

//UpdateFixture ...
func UpdateFixture(d *devices.Device) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Update Fixture")
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var updateFixture fixtures.BasicFixture
		json.Unmarshal(reqBody, &updateFixture)

		index, fixture := d.BasicFixtures.Find(updateFixture.Name)
		if index == -1 {
			fmt.Fprintf(w, "ERROR: The Fixture Was Not Found. Try Creating")
		}
		d.BasicFixtures.Update(index, fixture)
		json.NewEncoder(w).Encode(d.BasicFixtures[index])
	}
}

//DeleteFixture ...
func DeleteFixture(d *devices.Device) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete Fixture")
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var deleteFixture fixtures.BasicFixture
		json.Unmarshal(reqBody, &deleteFixture)

		index, _ := d.BasicFixtures.Find(deleteFixture.Name)
		if index != -1 {
			d.BasicFixtures.Delete(index, true)
			fmt.Fprintf(w, "Deleted.")
		}
		fmt.Fprintf(w, "Not Found.")
	}
}
