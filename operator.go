package main

import (
	_ "bufio"  // The underscore _ means we import but don't use this package
	 "fmt"     // This package is used for input/output operations
	_ "os"     // Not used in this program (ignored with _)
	_ "strconv" // Not used in this program (ignored with _)
	_ "text/template/parse"  // _ Blank identifier ignoring the imported package
)

// The main function is where program execution begins
func main() {
    // ==========================================
    // VARIABLE DECLARATION AND INITIALIZATION
    // ==========================================
    
    // Short variable declaration with :=
    // Creates two integer variables a and b in one line
    a := 10  // a is assigned the value 10 (type inferred as int)
    b := 3   // b is assigned the value 3 (type inferred as int)
    
    // ==========================================
    // BASIC ARITHMETIC OPERATORS
    // ==========================================

    // ADDITION: The + operator adds two numbers
    sum := a + b  // 10 + 3 = 13
    fmt.Println("Sum:", sum) // Prints: Sum: 13

    // ==========================================
    // EXPLICIT VARIABLE DECLARATION
    // ==========================================
    
    // Using var for explicit type declaration
    // Sometimes useful when you want to be specific about types
    var num1 int = 14  // Explicitly declare num1 as int with value 14
    var num2 int = 5   // Explicitly declare num2 as int with value 5

    // SUBTRACTION: The - operator subtracts the second number from the first
    difference := num1 - num2  // 14 - 5 = 9
    fmt.Println("Difference:", difference) // Prints: Difference: 9

    // MULTIPLICATION: The * operator multiplies two numbers
    product := a * b  // 10 * 3 = 30
    fmt.Println("Product:", product) // Prints: Product: 30

    // ==========================================
    // INTEGER DIVISION (TRUNCATES DECIMAL PART)
    // ==========================================
    
    // DIVISION: The / operator divides first number by second
    // WARNING: With integers, this performs INTEGER DIVISION
    // The decimal part is truncated (chopped off, not rounded)
    quotient := a / b  // 10 / 3 = 3.333... but integer result is 3
    fmt.Println("Quotient:", quotient) // Prints: Quotient: 3

    // ==========================================
    // FLOATING-POINT DIVISION (PRECISE)
    // ==========================================
    
    // Using float64 for precise decimal division
    var dividend float64 = 14  // Declare as float64 for decimal numbers
    var divisor float64 = 5    // Declare as float64 for decimal numbers

    // This division keeps the decimal part
    preciseQuotient := dividend / divisor  // 14.0 / 5.0 = 2.8
    fmt.Println("Precise Quotient:", preciseQuotient) // Prints: Precise Quotient: 2.8

    // ==========================================
    // TYPE CONVERSION (TYPE CASTING)
    // ==========================================
    
    // Different numeric types cannot be used together directly
    var numm float64 = 12    // 64-bit floating point number
    var numm2 uint16 = 6     // 16-bit unsigned integer (only positive numbers)

    // We must convert uint16 to float64 before multiplication
    // Type conversion syntax: desiredType(variable)
    preciseProduct := numm * float64(numm2) // Convert numm2 to float64, then multiply
    fmt.Println("Precise Product:", preciseProduct) // Prints: Precise Product: 72

    // ==========================================
    // MODULUS OPERATOR (REMAINDER)
    // ==========================================
    
    // Declare two more integer variables
    var val1 int = 17
    var val2 int = 4

    // MODULUS: The % operator gives the remainder after division
    // 17 divided by 4 = 4 with remainder 1 (because 4*4=16, 17-16=1)
    remainder := val1 % val2  // 17 % 4 = 1
    fmt.Println("Remainder:", remainder) // Prints: Remainder: 1
    
    // ==========================================
    // ADDITIONAL EXAMPLES AND EXPLANATIONS
    // ==========================================
    
    fmt.Println("\n=== EXTRA EXAMPLES ===")
    
    // Example 1: Why integer division truncates
    fmt.Println("\nInteger Division Examples:")
    fmt.Println("7 / 2 =", 7/2)   // Result: 3 (not 3.5)
    fmt.Println("9 / 4 =", 9/4)   // Result: 2 (not 2.25)
    
    // Example 2: Using parentheses to control order
    fmt.Println("\nOrder of Operations:")
    result1 := 2 + 3 * 4     // Multiplication first: 3*4=12, then 2+12=14
    result2 := (2 + 3) * 4   // Parentheses first: 2+3=5, then 5*4=20
    fmt.Println("2 + 3 * 4 =", result1)     // 14
    fmt.Println("(2 + 3) * 4 =", result2)   // 20
    
    // Example 3: Modulus with different numbers
    fmt.Println("\nModulus Examples:")
    fmt.Println("10 % 3 =", 10%3)   // 1 (10 divided by 3 is 3 with remainder 1)
    fmt.Println("15 % 4 =", 15%4)   // 3 (15 divided by 4 is 3 with remainder 3)
    fmt.Println("8 % 2 =", 8%2)     // 0 (8 is exactly divisible by 2)
    
    // Example 4: Checking even/odd using modulus
    number := 17
    if number % 2 == 0 {
        fmt.Printf("\n%d is even (divisible by 2)\n", number)
    } else {
        fmt.Printf("\n%d is odd (not divisible by 2)\n", number)
    }


//    Expected Output: 

//     Sum: 13
// Difference: 9
// Product: 30
// Quotient: 3
// Precise Quotient: 2.8
// Precise Product: 72
// Remainder: 1

// === EXTRA EXAMPLES ===

// Integer Division Examples:
// 7 / 2 = 3
// 9 / 4 = 2

// Order of Operations:
// 2 + 3 * 4 = 14
// (2 + 3) * 4 = 20

// Modulus Examples:
// 10 % 3 = 1
// 15 % 4 = 3
// 8 % 2 = 0

// 17 is odd (not divisible by 2)
}