package main

import "fmt"

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	fmt.Println("tv on")
	t.isRunning = true
}

func (t *tv) off() {
	fmt.Println("tv off")
	t.isRunning = false
}
