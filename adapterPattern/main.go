package main

import "fmt"

// the target interface
type UsbConnector interface {
	ConnectToUSBPort()
}

// the adaptee interface

type LighteningConnector interface {
	ConnectToLighteningPort()
}

// LightningCable is a concrete implementation of LightningConnector
type LightningCable struct{}

func (c *LightningCable) ConnectToLighteningPort() {
	fmt.Println("Lightning cable is connected to the device.")
}

// adapter
type LightningToUSBAdapter struct {
	lightningCable *LightningCable
}

func (a *LightningToUSBAdapter) ConnectToUSBPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.lightningCable.ConnectToLighteningPort()
}

func main() {
	lightningCable := &LightningCable{}
	adapter := &LightningToUSBAdapter{
		lightningCable: lightningCable,
	}

	// Client code just knows about the USB interface
	var usbPort UsbConnector = adapter
	usbPort.ConnectToUSBPort()
}
