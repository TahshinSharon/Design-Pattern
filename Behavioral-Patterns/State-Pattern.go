package main

import "fmt"

// State
type tvState interface {
	state()
}

// Concrete Implementation of State
type on struct{}

func (o *on) state() { // Implementing the Behaviour of ON State
	fmt.Println("Tv is On!")
}

type off struct{}

func (o *off) state() { // Implementing the Behaviour of Off State
	fmt.Println("Tv is off!")
}

// Context
type stateContext struct {
	currentTvState tvState
}

func getContext() *stateContext {
	return &stateContext{
		currentTvState: &off{},
	}
}
func (sc *stateContext) setState(state tvState) {
	sc.currentTvState = state
}

func (sc *stateContext) getState() {
	sc.currentTvState.state()
}
func main() {
	tvContext := getContext() // Default state is OFF
	tvContext.getState()      // Get The State As OFF
	tvContext.setState(&on{}) //Change the current state to ON
	tvContext.getState()      //Get the state as OFF
}
