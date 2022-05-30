package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(patient *patient) {
	if patient.paymentDone {
		return
	}

	fmt.Println("execute payment")
	patient.paymentDone = true
}

func (c *cashier) setNext(department department) {
	c.next = department
}
