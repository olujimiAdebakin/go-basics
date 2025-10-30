
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Birthday() {
	u.Age++ // modifies the original struct
}

func main() {
	u := User{Name: "Olujimi", Age: 24}
	fmt.Println("Before:", u)

	u.Birthday()
	fmt.Println("After:", u)

	x := 20

// Dereferencing Explained

// To dereference a pointer is to follow the address to get (or set) the actual value.
// Why? Because p points to the same memory as x.
// Changing *p changes the original variable.

p := &x  // p points to x

fmt.Println(*p) // read value → prints 20

*p = 50         // write value through pointer
fmt.Println(x)  // x is now 50

}
