# tinygo-expresslrs

A tinygo project to handle CRSF packets received by ExpressLRS modules.
This means that one can use ExpressLRS outside of flightcontrollers etc. and on hobby projects using a pi pico or any other UART capable microcontroller.

## Getting started

To get code highlighting and lsp support with zed, you can use the following job:

> [!WARNING]
> You cannot flash the device from a child process as this job sets environment variables which seem to be conflicting.
> Please start a new terminal to flash your device

```bash
just editor pico2 # Or other device
```

## References
