package main

import (
	"machine"
	"time"
)

const CRSF_BAUD = 420000

func main() {
	// 1. Configure UART0
	uart := machine.UART0
	err := uart.Configure(machine.UARTConfig{
		BaudRate: CRSF_BAUD,
		TX:       machine.GPIO0,
		RX:       machine.GPIO1,
	})
	if err != nil {
		println("UART Config Error:", err.Error())
	}

	println("CRSF Receiver Active - All 16 Channels")

	parser := NewCRSFParser()

	for {
		// Read available bytes
		b, err := uart.ReadByte()
		if err != nil {
			// No data, wait slightly to prevent CPU pinning if UART is empty
			time.Sleep(time.Microsecond * 10)
			continue
		}

		channels, err := parser.Feed(b)
		if err != nil {
			continue
		}

		if channels != nil {
			// Group 1: Joysticks (CH1-CH4)
			print("JOY: [")
			for i := 0; i < 4; i++ {
				print(channels[i])
				if i < 3 {
					print(" ")
				}
			}

			// Group 2: Aux Channels (CH5-CH16)
			print("] AUX: [")
			for i := 4; i < 16; i++ {
				print(channels[i])
				if i < 15 {
					print(" ")
				}
			}
			println("]")
		}
	}
}
