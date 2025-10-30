package main

import "fmt"

/*
RANGE DEFINITION:
================
range is a Go keyword used in for loops to iterate over:
- Arrays and Slices (get index + value)
- Maps (get key + value) 
- Strings (get index + rune/character)
- Channels (get values)

Syntax: for index, value := range collection { }
*/

func main() {
    // ========== BASIC ARRAY ITERATION ==========
    var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    // Traditional for loop - manual indexing
    fmt.Println("=== Traditional For Loop ===")
    for i := 0; i < len(a); i++ {
        fmt.Println("Element at index", i, ":", a[i])
    }

    // Range-based for loop - cleaner and safer
    fmt.Println("\n=== Range-Based For Loop ===")
    for i, element := range a {
        fmt.Println("Index:", i, "Value:", element)
    }

    // ========== STRING SLICE EXAMPLES ==========
    
    // Candy example
    fmt.Println("\n=== Candy Box Example ===")
    candies := []string{"🍭", "🍫", "🍬", "🍎", "🍪"}
    fmt.Println("Let's look at all our candies:")
    for position, candy := range candies {
        fmt.Printf("Candy #%d is: %s\n", position, candy)
    }

    // Student names example - FIXED VARIABLE NAMES
    fmt.Println("\n=== Student Names Example ===")
    studentNames := []string{"lekan", "tolu", "tosin", "jummy", "fola"}
    // CORRECTED: index, name (not name, gender)
    for index, name := range studentNames {
        fmt.Printf("Student %d is: %s\n", index, name)
    }

    // ========== MORE PRACTICAL EXAMPLES ==========
    
    // Classroom seating
    fmt.Println("\n=== Classroom Seating ===")
    students := []string{"Alice", "Bob", "Charlie", "Diana"}
    fmt.Println("Roll call by seat number:")
    for seatNumber, student := range students {
        fmt.Printf("Seat %d: %s\n", seatNumber, student)
    }

    // Recipe steps with human-friendly numbering
    fmt.Println("\n=== Recipe Steps ===")
    steps := []string{"Mix flour", "Add eggs", "Bake", "Decorate"}
    fmt.Println("Cake recipe steps:")
    for stepNumber, step := range steps {
        fmt.Printf("Step %d: %s\n", stepNumber+1, step) // +1 for human counting
    }

    // ========== FIXED CODE + MORE EXAMPLES ==========
    
    // Counting candies - FIXED: position++ to position+1
    fmt.Println("\n=== Counting Candies ===")
    candiees := []string{"🍭", "🍫", "🍬", "🍎", "🍪"}
    fmt.Println("Counting how many candies I have:")
    for position := range candiees {
        fmt.Printf("I have at least %d candies!\n", position+1) // FIXED: position+1
    }

    // ========== ADDITIONAL RANGE EXAMPLES ==========
    
    // Using only index (ignoring value)
    fmt.Println("\n=== Using Only Index ===")
    colors := []string{"red", "green", "blue", "yellow"}
    for index := range colors {
        fmt.Printf("Color at position %d exists!\n", index)
    }

    // Using only value (ignoring index)
    fmt.Println("\n=== Using Only Value ===")
    fruits := []string{"apple", "banana", "cherry", "date"}
    for _, fruit := range fruits {
        fmt.Printf("I like %s!\n", fruit)
    }

    // Range with maps
    fmt.Println("\n=== Range with Maps ===")
    studentAges := map[string]int{
        "Alice":   23,
        "Bob":     25,
        "Charlie": 22,
    }
    for name, age := range studentAges {
        fmt.Printf("%s is %d years old\n", name, age)
    }

    // Range with strings (gets Unicode characters)
    fmt.Println("\n=== Range with Strings ===")
    greeting := "Hello"
    for position, char := range greeting {
        fmt.Printf("Position %d: character '%c' (Unicode: %U)\n", position, char, char)
    }

    // ========== ADVANCED EXAMPLES ==========
    
    // Finding specific elements
    fmt.Println("\n=== Finding Elements ===")
    numbers := []int{10, 20, 30, 40, 50}
    target := 30
    for index, value := range numbers {
        if value == target {
            fmt.Printf("Found %d at position %d!\n", target, index)
            break // Stop searching once found
        }
    }

    // Summing values
    fmt.Println("\n=== Summing Values ===")
    prices := []float64{1.99, 2.49, 0.99, 3.99}
    total := 0.0
    for _, price := range prices {
        total += price
    }
    fmt.Printf("Total price: $%.2f\n", total)

    // Nested ranges (2D slice)
    fmt.Println("\n=== 2D Matrix ===")
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    for rowIndex, row := range matrix {
        for colIndex, value := range row {
            fmt.Printf("matrix[%d][%d] = %d\n", rowIndex, colIndex, value)
        }
    }
}

// ========== KEY POINTS TO REMEMBER ==========
/*
1. range returns (index, value) for arrays/slices
2. range returns (key, value) for maps  
3. range returns (index, rune) for strings
4. Use _ to ignore values you don't need
5. The loop variable is REUSED each iteration (be careful with pointers!)
6. range is generally safer than manual indexing (no out-of-bounds errors)
7. The order in maps is random, in slices it's guaranteed


Output Preview:

=== Candy Box Example ===
Let's look at all our candies:
Candy #0 is: 🍭
Candy #1 is: 🍫
Candy #2 is: 🍬
...

=== Student Names Example ===
Student 0 is: lekan
Student 1 is: tolu
...

=== Counting Candies ===
Counting how many candies I have:
I have at least 1 candies!
I have at least 2 candies!
...
*/