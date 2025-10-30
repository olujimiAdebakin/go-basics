package main

import "fmt"

/*
	👉 INTERFACES IN GOLANG

	An interface in Go defines a set of method signatures (behaviors)
	that a type must implement. Any type that provides definitions
	for all methods in an interface is said to "satisfy" that interface.

	Interfaces enable polymorphism — allowing different types to be
	used interchangeably if they share the same behavior.
*/

// ------------------------------------------------------------
// Shape interface — defines geometric behaviors for any shape
// ------------------------------------------------------------
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ------------------------------------------------------------
// Rectangle struct — a concrete type implementing Shape
// ------------------------------------------------------------
type Rectangle struct {
	Width, Height float64
}

// Rectangle implements all methods in the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ------------------------------------------------------------
// Circle struct — another concrete type implementing Shape
// ------------------------------------------------------------
type Circle struct {
	Radius float64
}

// Circle implements the Shape interface
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// ------------------------------------------------------------
// PaymentProcessor interface — defines a common payment behavior
// ------------------------------------------------------------
type PaymentProcessor interface {
	Process(amount float64) string
}

// ------------------------------------------------------------
// Paystack implementation of PaymentProcessor
// ------------------------------------------------------------
type Paystack struct{}

func (p Paystack) Process(amount float64) string {
	return fmt.Sprintf("Processed ₦%.2f through Paystack", amount)
}

// ------------------------------------------------------------
// Flutterwave implementation of PaymentProcessor
// ------------------------------------------------------------
type Flutterwave struct{}

func (f Flutterwave) Process(amount float64) string {
	return fmt.Sprintf("Processed ₦%.2f through Flutterwave", amount)
}

// ------------------------------------------------------------
// MakePayment function — accepts any PaymentProcessor
// ------------------------------------------------------------
// This demonstrates polymorphism: different payment processors
// can be passed to the same function as long as they implement
// the PaymentProcessor interface.
func MakePayment(p PaymentProcessor, amount float64) {
	fmt.Println(p.Process(amount))
}

// ------------------------------------------------------------
// MAIN FUNCTION
// ------------------------------------------------------------
func main() {
	// --- Working with Shape Interface ---
	var s Shape // Declare a variable of type Shape

	// Assign Rectangle (implements Shape)
	s = Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())

	// Assign Circle (also implements Shape)
	s = Circle{Radius: 7}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())

	// --- Working with PaymentProcessor Interface ---
	paystack := Paystack{}
	flutterwave := Flutterwave{}

	// Both types satisfy PaymentProcessor interface
	MakePayment(paystack, 5000)
	MakePayment(flutterwave, 12000)
}

/*
	🧠 Summary:

	- `Shape` and `PaymentProcessor` are interfaces that define behaviors.
	- `Rectangle`, `Circle`, `Paystack`, and `Flutterwave` are concrete types
	  that implement those behaviors.
	- The `MakePayment` function shows polymorphism — it accepts any type
	  that satisfies the `PaymentProcessor` interface.

	This is Go's version of interface-based polymorphism — no inheritance,
	no class hierarchies — just behavior contracts.
*/
