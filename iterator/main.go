package main

import (
	"fmt"
)

type Person struct {
	habits []string
}

func (p *Person) listHabits() {
	for _, h := range p.habits {
		fmt.Println(h)
	}
}

func main() {
	p := &Person{
		habits: []string{"basketball", "baseball", "jogging", "swimming"},
	}
	p.listHabits()
}
