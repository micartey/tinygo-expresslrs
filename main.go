package main

import (
	"machine"

	"github.com/soypat/cyw43439"
)

const (
	LED = machine.LED

	// UART_ANTENNA_TX = machine.GPIO0
	// UART_ANTENNA_RX = machine.GPIO1
)

func main() {
	// 1. Initialize the WiFi chip (CYW43439)
	// On Pico 2 W, this chip controls the LED.
	// We use the default SPI pins defined by the board target.
	dev := cyw43439.NewPicoWDevice()
	cfg := cyw43439.DefaultWifiConfig()
	err := dev.Init(cfg)
	if err != nil {
		println("Could not configure WiFi device:", err.Error())
		return
	}

	// 2. The LED is 'GPIO 0' on the *WiFi chip* (not the RP2350)
	dev.GPIOSet(0, true)

	println("WiFi initialized. Blinking LED...")

	select {}
}
