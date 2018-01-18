//shows how to watch for new devices and list them
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/emitter"
	"github.com/muka/go-bluetooth/linux"
)

const logLevel = log.DebugLevel
const adapterID = "hci0"

func main() {

	log.SetLevel(logLevel)

	//clean up connection on exit
	defer api.Exit()

	log.Debugf("Reset bluetooth device")
	a := linux.NewBtMgmt(adapterID)
	err := a.Reset()
	if err != nil {
		panic(err)
	}

	err = api.ClearDevices()
	if err != nil {
		panic(err)
	}

	err = discoverDevices(adapterID)
	if err != nil {
		panic(err)
	}

	select {}
}

func discoverDevices(adapterID string) error {

	err := api.StartDiscovery()
	if err != nil {
		return err
	}

	log.Debugf("Started discovery")
	api.On("discovery", emitter.NewCallback(func(ev emitter.Event) {
		discoveryEvent := ev.GetData().(api.DiscoveredDeviceEvent)
		dev := discoveryEvent.Device
		showDeviceInfo(dev)
	}))

	return nil
}

func showDeviceInfo(dev *api.Device) {
	if dev == nil {
		return
	}
	props, err := dev.GetProperties()
	if err != nil {
		log.Errorf("%s: Failed to get properties: %s", dev.Path, err.Error())
		return
	}
	log.Infof("name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)
}
