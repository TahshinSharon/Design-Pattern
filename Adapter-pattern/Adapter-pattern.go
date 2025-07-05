package main

import "fmt"

// Target
type mobile interface {
	chargeAppleMobile()
}

// Concrete Prototype Implementation
type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Charging Apple Mobile")
}

// Adaptee
type android struct{}

func (ad *android) chargeAndroidMobile() {
	fmt.Println("Charging Android Mobile")
}

// Adapter

type androidAdapter struct {
	android *android
}

func (ad *androidAdapter) chargeAppleMobile() {
	ad.android.chargeAndroidMobile()
}

// Client
type client struct{}

func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleMobile()
}
func main() {
	// Initial Requirment
	apple := &apple{}

	client := &client{}

	client.chargeMobile(apple)

	// Extended Requirement i.e Charge Android Mobile

	android := &android{}
	androidAdapter := &androidAdapter{
		android: android,
	}
	client.chargeMobile(androidAdapter)
}

/*
What Is  Adapter Pattern
In Golang, the Adapter Pattern is used to make a type compatible with an interface it doesn’t implement directly. It helps convert one type into another by wrapping it with an adapter struct that implements the target interface.

Go doesn’t use inheritance — instead, it relies heavily on interfaces and composition, making the adapter pattern a perfect fit.

Why use the Adapter Pattern in Go?
	1.	Interface compatibility: To allow a type with a different method signature to satisfy an expected interface.
	2.	Legacy code integration: To use older or third-party code without modifying it.
	3.	Plug-and-play architecture: To make different components work together cleanly by adapting their interfaces.

*/
