package main

import (
	"machine"
)

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
