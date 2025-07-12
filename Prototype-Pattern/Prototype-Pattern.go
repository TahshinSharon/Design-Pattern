package main

import (
	"fmt"
	"math"
)

type ShapeType int

const (
	CircleType ShapeType = 1
	SquareType ShapeType = 2
)

type Shape interface {
	GetId() ShapeType // Get The Shape Id
	PrintTypeProp()   // Used For Printing Property Values Of The Shape
	Clone() Shape     // Used For Getting DeepCopy
}

type Circle struct {
	Id            ShapeType
	Radius        float32
	Diameter      float32
	Circumference float32
}

func NewCircle(Radius float32, Diameter float32, Circumference float32) Circle {
	return Circle{CircleType, Radius, Diameter, Circumference}
}

func (c Circle) GetId() ShapeType {
	return c.Id
}
func (c Circle) Clone() Shape { // Prototyping
	return NewCircle(c.Radius, c.Diameter, c.Circumference)
}
func (c Circle) PrintTypeProp() {
	fmt.Println("Circle Properties Radius:", c.Radius, " Diameter:", c.Diameter, "Circumference:", c.Circumference, " Area:", math.Pi*c.Radius)
}

type Square struct {
	Id     ShapeType
	Length float32
}

func NewSquare(length float32) Square {
	return Square{SquareType, length}
}
func (s Square) GetId() ShapeType {
	return s.Id
}
func (s Square) Clone() Shape { // Prototyping
	return NewSquare(s.Length)
}
func (s Square) PrintTypeProp() {
	fmt.Println("Square Properties Length:", s.Length, " Area:", s.Length*s.Length)
}

var RegistryList = make(map[int]Shape)

func loadToRegistry() {
	circle := NewCircle(50, 40, 15)
	RegistryList[int(circle.GetId())] = circle //Adding circle to Registry

	square := NewSquare(20)

	RegistryList[int(square.GetId())] = square //Adding Square To Registry
}

func main() {
	loadToRegistry() //Load New Objects data to Registry

	square := RegistryList[int(SquareType)]
	sq, ok := square.(Square)
	if ok {
		fmt.Printf("Old ")
		sq.PrintTypeProp()
		NewSquare := sq.Clone() //Prototype
		fmt.Printf("Cloned ")
		NewSquare.PrintTypeProp()
	}

	circle := RegistryList[int(CircleType)]
	cr, ok := circle.(Circle) //Type Assertion
	if ok {
		fmt.Printf("Old ")
		cr.PrintTypeProp()
		NewCircle := cr.Clone().(Circle) // Prototype i.e Cloning existing object
		NewCircle.Radius = 15            // Changin property of the prototype without effecting the original object
		fmt.Printf("Cloned ")
		NewCircle.PrintTypeProp()
	}
}
