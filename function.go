package main

import (
	"fmt"
	"strings"
)

/*
A function is a reusable block of code that performs a specific task. 
Think of it like a mini-program within your main program that you can call whenever you need it.

BASIC SYNTAX:
func functionName(parameters) returnType {
    // code to execute
    return value
}
*/

// ========== BASIC FUNCTION DEFINITIONS ==========

// Function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello, World! 👋")
}

// Function with one parameter and no return value
func greetPerson(name string) {
	fmt.Printf("Hello, %s! 🎉 \n", name)
}

// Function with parameters and return value
func addNumbers(a int, b int) int {
	return a + b
}

// Function with multiple parameters of same type
func multiply(x int, y int) int {
	return x * y
}

// Function with multiple return values
func calculate(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

// Single return value
func getCircleArea(radius float64) float64 {
	return 3.14159 * radius * radius
}

// Multiple return values with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func getDimensions(length, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return - returns area and perimeter
}

// Function returning a boolean
func isEven(number int) bool {
	return number%2 == 0
}

// Function returning a string
func getGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// ========== CALCULATOR FUNCTIONS ==========

func calculator() {
	fmt.Println("\n=== CALCULATOR FUNCTIONS ===")
	
	fmt.Printf("15 + 27 = %d\n", add(15, 27))
	fmt.Printf("15 - 27 = %d\n", subtract(15, 27))
	fmt.Printf("15 × 27 = %d\n", multiply(15, 27))
	
	result, err := divide(15, 3)
	if err == nil {
		fmt.Printf("15 ÷ 3 = %.1f\n", result)
	}
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// ========== STRING MANIPULATION FUNCTIONS ==========

func stringFunctions() {
	fmt.Println("\n=== STRING FUNCTIONS ===")
	
	name := "  John Doe  "
	fmt.Printf("Original: '%s'\n", name)
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(name))
	fmt.Printf("Uppercase: '%s'\n", strings.ToUpper(name))
	fmt.Printf("Reversed: '%s'\n", reverseString("Hello"))
	
	fmt.Printf("Is 'racecar' palindrome? %t\n", isPalindrome("racecar"))
	fmt.Printf("Is 'hello' palindrome? %t\n", isPalindrome("hello"))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	return s == reverseString(s)
}

// ========== PARAMETER EXAMPLES ==========

func parameterExamples() {
	fmt.Println("\n=== PARAMETER EXAMPLES ===")
	
	introduce("Alice", 25)
	introduce("", 30) // Uses default name
	
	// Variadic function calls
	fmt.Printf("Sum: %d\n", sumNumbers(1, 2, 3))
	fmt.Printf("Sum: %d\n", sumNumbers(10, 20, 30, 40, 50))
	
	// Using slice with variadic function
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of slice: %d\n", sumNumbers(numbers...))
	
	// Slice parameter
	grades := []float64{85.5, 92.0, 78.5, 96.0}
	fmt.Printf("Average grade: %.2f\n", averageGrades(grades))
	
	// Map parameter
	studentGrades := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	printStudentGrades(studentGrades)
}

// Function with default-like behavior
func introduce(name string, age int) {
	if name == "" {
		name = "Anonymous"
	}
	fmt.Printf("Hello, I'm %s and I'm %d years old\n", name, age)
}

// Variadic function - accepts any number of arguments
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function with slice parameter
func averageGrades(grades []float64) float64 {
	if len(grades) == 0 {
		return 0
	}
	total := 0.0
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))
}

// Function with map parameter
func printStudentGrades(grades map[string]int) {
	fmt.Println("Student Grades:")
	for student, grade := range grades {
		fmt.Printf("  %s: %d\n", student, grade)
	}
}

// ========== ADVANCED CONCEPTS ==========

// Higher-order function - function that takes another function as parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Function that returns another function
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Recursive function - function that calls itself
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Closure - function that captures variables from its surrounding scope
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func advancedExamples() {
	fmt.Println("\n=== ADVANCED FUNCTION CONCEPTS ===")
	
	// Higher-order function
	add := func(a, b int) int { return a + b }
	multiplyFunc := func(a, b int) int { return a * b }
	
	fmt.Printf("10 + 5 = %d\n", applyOperation(10, 5, add))
	fmt.Printf("10 × 5 = %d\n", applyOperation(10, 5, multiplyFunc))
	
	// Function returning function
	double := multiplier(2)
	triple := multiplier(3)
	
	fmt.Printf("Double 5: %d\n", double(5))
	fmt.Printf("Triple 5: %d\n", triple(5))
	
	// Recursive function
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Factorial of 7: %d\n", factorial(7))
	
	// Closure
	myCounter := counter()
	fmt.Println("Counter:", myCounter()) // 1
	fmt.Println("Counter:", myCounter()) // 2
	fmt.Println("Counter:", myCounter()) // 3
	
	anotherCounter := counter()
	fmt.Println("Another counter:", anotherCounter()) // 1
}

// ========== REAL-WORLD EXAMPLES ==========

// Banking system functions
type Account struct {
	balance float64
	owner   string
}

func createAccount(owner string, initialDeposit float64) Account {
	return Account{
		balance: initialDeposit,
		owner:   owner,
	}
}

func (a *Account) deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, a.balance)
}

func (a *Account) withdraw(amount float64) bool {
	if amount > a.balance {
		fmt.Printf("Insufficient funds. Balance: $%.2f\n", a.balance)
		return false
	}
	a.balance -= amount
	fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, a.balance)
	return true
}

func (a Account) getBalance() float64 {
	return a.balance
}

// E-commerce functions
func calculateTotalPrice(price float64, quantity int, discount float64) float64 {
	subtotal := price * float64(quantity)
	discountAmount := subtotal * (discount / 100)
	return subtotal - discountAmount
}

func applyTax(amount float64, taxRate float64) float64 {
	tax := amount * (taxRate / 100)
	return amount + tax
}

func realWorldExamples() {
	fmt.Println("\n=== REAL-WORLD EXAMPLES ===")
	
	// Banking example
	fmt.Println("🏦 BANKING SYSTEM")
	account := createAccount("Alice", 1000.0)
	account.deposit(500.0)
	account.withdraw(200.0)
	account.withdraw(2000.0) // Should fail
	
	// E-commerce example
	fmt.Println("\n🛒 E-COMMERCE SYSTEM")
	price := 25.50
	quantity := 3
	discount := 10.0 // 10% discount
	taxRate := 8.0   // 8% tax
	
	subtotal := calculateTotalPrice(price, quantity, discount)
	total := applyTax(subtotal, taxRate)
	
	fmt.Printf("Price per item: $%.2f\n", price)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Discount: %.1f%%\n", discount)
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax (%.1f%%): $%.2f\n", taxRate, total-subtotal)
	fmt.Printf("Total: $%.2f\n", total)
}

// ========== MAIN FUNCTION ==========

func mainBasicExamples() {
	fmt.Println("=== BASIC FUNCTION EXAMPLES ===")

	greetPerson("jimi") // Hello, Jimi 🎉
	sayHello()          // Hello, World 👋

	result := addNumbers(5, 3) // Returns 8
	fmt.Println("5 + 3 =", result)

	product := multiply(4, 7) // Returns 28
	fmt.Println("4 × 7 =", product)

	sum, prod := calculate(6, 8) // Returns 14 and 48
	fmt.Printf("6+8=%d, 6×8=%d\n", sum, prod)

	// Return value examples
	fmt.Println("\n=== RETURN VALUE EXAMPLES ===")

	area := getCircleArea(5.0)
	fmt.Printf("Circle area with radius 5: %.2f\n", area)

	divisionResult, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.1f\n", divisionResult)
	}

	// Try division by zero (ignore result, only check error)
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	rectArea, perimeter := getDimensions(10, 5)
	fmt.Printf("Rectangle: area=%d, perimeter=%d\n", rectArea, perimeter)

	fmt.Printf("Is 7 even? %t\n", isEven(7))
	fmt.Printf("Is 8 even? %t\n", isEven(8))

	fmt.Printf("Score 85 gets grade: %s\n", getGrade(85))
	fmt.Printf("Score 45 gets grade: %s\n", getGrade(45))
}

func main() {
	fmt.Println("🎯 FUNCTIONS IN GO - COMPLETE GUIDE")
	fmt.Println("===================================")
	
	// Run all examples
	mainBasicExamples()
	calculator()
	stringFunctions()
	parameterExamples()
	advancedExamples()
	realWorldExamples()
	
	fmt.Println("\n=== FUNCTION GUIDE COMPLETE ===")
}