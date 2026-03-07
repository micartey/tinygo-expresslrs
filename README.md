# tinygo-expresslrs

A tinygo project to handle CRSF packets received by ExpressLRS modules.
This means that one can use ExpressLRS outside of flightcontrollers etc. and on hobby projects using a pi pico or any other UART capable microcontroller.

## Getting started

To get code highlighting and lsp support with zed, you can use the following job:

> [!WARNING]
> You cannot flash the device from a child process as this job sets environment variables which seem to be conflicting.
> Please start a new terminal to flash your device

```bash
just zed-editor pico2 # Or other device
```

### Usage

Add the module to your TinyGo project:

```bash
go get github.com/micartey/tinygo-expresslrs
```

Then import and use the CRSF parser in your code:

```go
package main

import (
	"machine"
	"time"

	crsf "github.com/micartey/tinygo-expresslrs"
)

func main() {
	uart := machine.UART0
	uart.Configure(machine.UARTConfig{
		BaudRate: 420000,
		TX:       machine.GPIO0,
		RX:       machine.GPIO1,
	})

	parser := crsf.NewCRSFParser()

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

		// channels[0] through channels[15]
	}
}
```

## Pinout

The project uses by default **UART0** _(TX: GP0, RX: GP1)_.
You can take te power from ping 38 and 40.
Using Pin 40 will allow you to draw 5V which is generally what you should provide to an ExpressLRS Receiver.

![img](https://www.raspberrypi.com/documentation/microcontrollers/images/pico2w-pinout.svg)

_(VSYS is used to supply 5V to the Pico if you want to operate it externally - this is very risky and can easily fry your controller)_
