package main

import (
	"machine"
	"time"
)

const CRSF_BAUD = 420000

// Channel values extracted to variables outside of main
var (
	CH1, CH2, CH3, CH4     uint16
	CH5, CH6, CH7, CH8     uint16
	CH9, CH10, CH11, CH12  uint16
	CH13, CH14, CH15, CH16 uint16

	Connected     bool
	LastFrameTime time.Time
)

func main() {
	// Configure UART0
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

	// Start the receiver in its own goroutine
	go runReceiver(uart)

	// Main loop can now perform other tasks while channels update in background
	for {
		if !Connected {
			println("Receiver Disconnected...")
		} else {
			if time.Since(LastFrameTime) > 500*time.Millisecond {
				Connected = false
			}
			print("[", CH3, " ", CH4, " - ", CH1, " ", CH2, "] ", CH10, "\r\n")
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func runReceiver(uart *machine.UART) {
	parser := NewCRSFParser()

	for {
		b, err := uart.ReadByte()
		if err != nil {
			time.Sleep(time.Microsecond * 10)
			continue
		}

		channels, err := parser.Feed(b)
		if err != nil || channels == nil {
			continue
		}

		// Update connection status
		Connected = true
		LastFrameTime = time.Now()

		// Update global variables
		CH1, CH2, CH3, CH4 = channels[0], channels[1], channels[2], channels[3]
		CH5, CH6, CH7, CH8 = channels[4], channels[5], channels[6], channels[7]
		CH9, CH10, CH11, CH12 = channels[8], channels[9], channels[10], channels[11]
		CH13, CH14, CH15, CH16 = channels[12], channels[13], channels[14], channels[15]
	}
}
