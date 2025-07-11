package main

import (
	"fmt"
)

// Subject
type Subject struct { // Iphone
	observers []Observer // List of Observer To Notify
	state     int
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) ChangeState(state int) {
	s.state = state

	//Notify All Observers
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

//Observer Interface

type Observer interface {
	Update(state int)
}
type ConcreateObserver struct {
	id int
}

func (o *ConcreateObserver) Update(state int) {
	fmt.Printf("Observer %d Received State Update: %d\n", o.id, state)
}
func main() {
	//Client
	sub := &Subject{}
	// Crate Observer
	ob1 := &ConcreateObserver{1}
	ob2 := &ConcreateObserver{2}
	//Attach Observer
	sub.Attach(ob1)
	sub.Attach(ob2)
	//Change The State Of Subject
	sub.ChangeState(1) //Notify All The Observers
}
