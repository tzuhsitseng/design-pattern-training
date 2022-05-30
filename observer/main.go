package main

import "fmt"

type observer interface {
	getID() int
	notify(msg string)
}

type myObserver struct {
	id int
}

func (m myObserver) getID() int {
	return m.id
}

func (m myObserver) notify(msg string) {
	fmt.Println(m.id, "got msg:", msg)
}

type pubsub struct {
	subscribers map[string]map[int]observer
}

func NewPubSub() *pubsub {
	return &pubsub{subscribers: map[string]map[int]observer{}}
}

func (p *pubsub) subscribe(topic string, ob observer) {
	if _, ok := p.subscribers[topic]; ok {
		p.subscribers[topic][ob.getID()] = ob
	} else {
		p.subscribers[topic] = map[int]observer{ob.getID(): ob}
	}
}

func (p *pubsub) publish(topic, msg string) {
	for _, s := range p.subscribers[topic] {
		s.notify(msg)
	}
}

func main() {
	ps := NewPubSub()
	ob1 := myObserver{id: 1}
	ob2 := myObserver{id: 2}
	ps.subscribe("A", ob1)
	ps.subscribe("A", ob2)
	ps.subscribe("B", ob1)
	ps.publish("A", "Hi everyone")
	ps.publish("B", "Good bye")
}
