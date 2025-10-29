package main

import "fmt"

func main() {
	// ==========================================
	// DATA TYPES IN GOLANG
	// ==========================================

	// STRING - sequence of characters
	var name string 
	name = "lekan"    // Initial assignment
	name = "Bull"     // Reassignment (variables can be updated)
	fmt.Println("Name:", name)

	// INTEGER TYPES - signed and unsigned integers of various sizes
	var number int16 = 4000          // 16-bit signed integer (-32768 to 32767)
	var age int = 25                 // Platform-dependent (32 or 64 bits)
	var bigNumber int64 = 1234567890 // 64-bit signed integer
	var smallNumber int8 = 127       // 8-bit signed integer (-128 to 127)
	
	fmt.Println("Number:", number)
	fmt.Println("Age:", age)
	fmt.Println("Big Number:", bigNumber)
	fmt.Println("Small Number:", smallNumber)

	// UNSIGNED INTEGER TYPES - only positive numbers
	var unsignedNumber uint = 500    // Only positive values
	var byteValue byte = 255         // Alias for uint8 (0 to 255)
	var unsignedBig uint64 = 18446744073709551615 // Maximum uint64 value
	
	fmt.Println("Unsigned Number:", unsignedNumber)
	fmt.Println("Byte Value:", byteValue)
	fmt.Println("Unsigned Number:", unsignedBig)

	// FLOATING POINT TYPES - decimal numbers
	var price float32 = 19.99        // 32-bit floating point
	var salary float64 = 55000.75    // 64-bit floating point (more precision)
	var temperature = 36.6           // Type inferred as float64
	
	fmt.Println("Price:", price)
	fmt.Println("Salary:", salary)
	fmt.Println("Temperature:", temperature)

	// BOOLEAN - true/false values
	var isActive bool = true
	var isCompleted = false          // Type inference
	var isLoggedIn bool              // Default value is false
	
	fmt.Println("Is Active:", isActive)
	fmt.Println("Is Completed:", isCompleted)
	fmt.Println("Is Logged In:", isLoggedIn)

	// COMPLEX NUMBERS - real and imaginary parts
	var complexNum complex64 = 3 + 4i
	var bigComplex complex128 = complex(5, 10) // 5 + 10i
	
	fmt.Println("Complex Number:", complexNum)
	fmt.Println("Big Complex:", bigComplex)

	// ARRAYS - fixed-size collections
	var numbers [3]int = [3]int{1, 2, 3}        // Explicit declaration
	var fruits = [2]string{"Apple", "Banana"}   // Type inference
	var emptyArray [5]int                        // All elements initialized to 0
	
	fmt.Println("Numbers Array:", numbers)
	fmt.Println("Fruits Array:", fruits)
	fmt.Println("Empty Array:", emptyArray)

	// SLICES - dynamic arrays (more flexible than arrays)
	var scores []int = []int{85, 90, 78}        // Slice literal
	var names []string                           // Nil slice
	names = append(names, "Alice", "Bob")       // Adding elements
	
	fmt.Println("Scores Slice:", scores)
	fmt.Println("Names Slice:", names)

	// MAPS - key-value pairs
	var person map[string]string = map[string]string{
		"name": "John",
		"city": "New York",
	}
	var ages map[string]int = make(map[string]int) // Using make function
	ages["Alice"] = 25
	ages["Bob"] = 30
	
	fmt.Println("Person Map:", person)
	fmt.Println("Ages Map:", ages)

	// STRUCT - custom data type
	type Employee struct {
		ID     int
		Name   string
		Salary float64
	}
	
	var emp Employee = Employee{
		ID:     1,
		Name:   "Jane Doe",
		Salary: 50000.0,
	}
	
	fmt.Println("Employee Struct:", emp)

	// POINTERS - store memory addresses
	var x int = 42
	var ptr *int = &x  // & gets the address, * declares pointer type
	
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)
	fmt.Println("Value through pointer:", *ptr) // * dereferences pointer

	// TYPE ALIASES
	type ID int
	type Currency float64
	
	var userID ID = 12345
	var amount Currency = 99.99
	
	fmt.Println("User ID:", userID)
	fmt.Println("Amount:", amount)

	// ZERO VALUES - default values for uninitialized variables
	var defaultInt int          // 0
	var defaultFloat float64    // 0.0
	var defaultString string    // ""
	var defaultBool bool        // false
	var defaultSlice []int      // nil
	
	fmt.Println("Default Int:", defaultInt)
	fmt.Println("Default Float:", defaultFloat)
	fmt.Println("Default String:", "'" + defaultString + "'")
	fmt.Println("Default Bool:", defaultBool)
	fmt.Println("Default Slice:", defaultSlice)

	// ==========================================
	// TYPE CONVERSION EXAMPLES
	// ==========================================
	var a int = 42
	var b float64 = float64(a)  // Explicit conversion required
	var c int = int(b)
	
	fmt.Println("Type Conversion - Int to Float:", b)
	fmt.Println("Type Conversion - Float to Int:", c)

	// ==========================================
	// MULTIPLE VARIABLE DECLARATION
	// ==========================================
	var (
		firstName = "John"
		lastName  = "Doe"
		userAge   = 30
	)
	
	fmt.Println("Multiple Variables:", firstName, lastName, userAge)

	fmt.Println("Hello, World! - All examples completed!")
}