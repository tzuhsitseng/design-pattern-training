package main

type department interface {
	execute(patient *patient)
	setNext(department department)
}
