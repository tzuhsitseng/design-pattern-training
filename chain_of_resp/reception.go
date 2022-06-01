package main

import "fmt"

type reception struct {
	next department
}

func (r *reception) execute(patient *patient) {
	if patient.registrationDone {
		r.next.execute(patient)
		return
	}

	fmt.Println("execute registration")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *reception) setNext(department department) {
	r.next = department
}
