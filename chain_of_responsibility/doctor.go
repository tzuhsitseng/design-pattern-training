package main

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(patient *patient) {
	if patient.doctorCheckDone {
		d.next.execute(patient)
		return
	}

	fmt.Println("execute doctor checking")
	patient.doctorCheckDone = true
	d.next.execute(patient)
}

func (d *doctor) setNext(department department) {
	d.next = department
}
