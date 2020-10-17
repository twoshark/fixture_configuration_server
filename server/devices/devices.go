package devices

import (
	"encoding/json"
	"io/ioutil"

	"github.com/twoshark/fixture_configuration_server/server/fixtures"
)

//Devices ...
type Devices struct {
	Name    string
	Devices []Device
}

//Device ...
type Device struct {
	Name          string                 `json:"name"`
	BasicFixtures fixtures.BasicFixtures `json:"fixtures"`
}

//NewDevices ...
func NewDevices(name string, devices ...Device) Devices {
	return Devices{
		Name:    name,
		Devices: devices,
	}
}

//Add ...
func (d *Devices) Add(newDevice Device) {
	d.Devices = append(d.Devices, newDevice)
}

//Find ...
func (d Devices) Find(name string) (int, Device) {
	for index, device := range d.Devices {
		if device.Name == name {
			return index, device
		}
	}
	return -1, Device{}
}

//Update ...
func (d *Devices) Update(index int, update Device) {
	d.Devices[index] = update
}

//Delete ...
func (d *Devices) Delete(index int, preserveOrder bool) {
	installs := d.Devices
	copy(installs[index:], installs[index+1:]) // Shift a[i+1:] left one index.
	installs[len(installs)-1] = Device{}       // Erase last element (write zero value).
	installs = installs[:len(installs)-1]      // Truncate slice.
	d.Devices = installs
}

//Save ...
func (d Devices) Save() {
	json, err := json.Marshal(d)
	if err != nil {
		_ = ioutil.WriteFile("_"+d.Name+".json", json, 0644)
	}

}
