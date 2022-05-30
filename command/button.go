package main

import "fmt"

type button struct {
	cmd command
}

func (b *button) press() {
	fmt.Println("press btn")
	b.cmd.execute()
}
