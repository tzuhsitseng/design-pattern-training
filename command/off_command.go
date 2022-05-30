package main

import "fmt"

type offCommand struct {
	device device
}

func (c offCommand) execute() {
	fmt.Println("do off command")
	c.device.off()
}
