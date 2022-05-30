package main

import "fmt"

type onCommand struct {
	device device
}

func (c onCommand) execute() {
	fmt.Println("do on command")
	c.device.on()
}
