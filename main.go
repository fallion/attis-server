package main

import iocontroller "github.com/fallion/attis-server/internal/io-controller"

func main() {
	service := iocontroller.Service{}

	service.TurnOnLED()
}
