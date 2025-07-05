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
