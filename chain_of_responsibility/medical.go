package main

import "fmt"

type medical struct {
	next department
}

func (m *medical) execute(patient *patient) {
	if patient.medicineDone {
		m.next.execute(patient)
		return
	}

	fmt.Println("execute medical")
	patient.medicineDone = true
	m.next.execute(patient)
}

func (m *medical) setNext(department department) {
	m.next = department
}
