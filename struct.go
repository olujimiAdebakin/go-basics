package main

import "fmt"

// Define a struct type
type User struct {
	Name  string
	Email string
	Age   int
}

type Person struct {
	Name  string
	Email string
}


// Attach method to struct
func (u User) Greet() {
	fmt.Printf("Hello, %s! Your email is %s\n", u.Name, u.Email)
}

// Pointer receiver method that updates field
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	// Create struct instance directly
	user1 := User{Name: "Olujimi", Email: "olu@example.com", Age: 25}

	// Or use the "dot" syntax to assign fields
	var user2 User
	user2.Name = "Ada"
	user2.Email = "ada@example.com"
	user2.Age = 30

	// Print struct values
	fmt.Println("User1:", user1)
	fmt.Println("User2:", user2)

	// Access individual fields
	fmt.Println("User1 name:", user1.Name)

	person := Person{Name: "Olujimi", Email: "olu@example.com"}

	person.Greet()

	person.UpdateEmail("new@paynode.com")

	person.Greet()

}
