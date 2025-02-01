package main

import (
	"machine"
)

var period uint64 = 1e9 / 500

func main() {
	button := machine.D2
	led := machine.LED

	button.Configure(machine.PinConfig{
		Mode: machine.PinInput,
	})

	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	for {
		if button.Get() {
			led.High()
		} else {
			led.Low()
		}
	}
}
