package main

import (
	_ "bufio"  // The underscore _ means we import but don't use this package
	 "fmt"     // This package is used for input/output operations
	_ "os"     // Not used in this program (ignored with _)
	_ "strconv" // Not used in this program (ignored with _)
	_ "text/template/parse"  // _ Blank identifier ignoring the imported package
)

func main() {
    fmt.Println("=== LOOPS IN GOLANG DEMONSTRATION ===")
    
    // ==========================================
    // 1. BASIC FOR LOOP (Traditional C-style)
    // ==========================================
    
    fmt.Println("\n--- 1. BASIC FOR LOOP ---")

    	

	// Example 1: Basic for loop
	fmt.Println("\n--- BASIC FOR LOOP ---")
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}


	x := 4

 	for x < 5 {
		fmt.Println("value of x:", x)
		x++
	}
    
    // Traditional for loop with initialization, condition, and post statement
    for i := 0; i < 5; i++ {
        fmt.Printf("Basic loop - Iteration %d\n", i)
    }
    
    // ==========================================
    // 2. WHILE-LIKE LOOP (Go only has 'for' loop)
    // ==========================================
    
    fmt.Println("\n--- 2. WHILE-STYLE LOOP ---")
    
    // While-like loop: only condition is specified
    count := 0
    for count < 3 {
        fmt.Printf("While-style - Count: %d\n", count)
        count++ // Don't forget to increment or you'll get infinite loop!
    }
    
    // ==========================================
    // 3. INFINITE LOOP (with break)
    // ==========================================
    
    fmt.Println("\n--- 3. INFINITE LOOP WITH BREAK ---")
    
    counter := 0
    for {
        fmt.Printf("Infinite loop - Counter: %d\n", counter)
        counter++
        
        // Condition to break out of infinite loop
        if counter >= 3 {
            fmt.Println("Breaking out of infinite loop!")
            break // Exit the loop immediately
        }
    }
    
    // ==========================================
    // 4. LOOPING THROUGH ARRAYS/SLICES
    // ==========================================
    
    fmt.Println("\n--- 4. LOOPING THROUGH SLICES ---")
    
    // Create a slice of strings
    fruits := []string{"Apple", "Banana", "Orange", "Grape", "Mango"}
    
    // Method 1: Traditional index-based loop
    fmt.Println("Method 1 - Index-based loop:")
    for i := 0; i < len(fruits); i++ {
        fmt.Printf("  Index %d: %s\n", i, fruits[i])
    }
    
    // Method 2: Range loop (index and value)
    fmt.Println("Method 2 - Range loop (index & value):")
    for index, fruit := range fruits {
        fmt.Printf("  Index %d: %s\n", index, fruit)
    }
    
    // Method 3: Range loop (value only)
    fmt.Println("Method 3 - Range loop (value only):")
    for _, fruit := range fruits { // _ ignores the index
        fmt.Printf("  Fruit: %s\n", fruit)
    }
    
    // Method 4: Range loop (index only)
    fmt.Println("Method 4 - Range loop (index only):")
    for index := range fruits { // ignores the value
        fmt.Printf("  Index: %d\n", index)
    }
    
    // ==========================================
    // 5. LOOPING THROUGH MAPS
    // ==========================================
    
    fmt.Println("\n--- 5. LOOPING THROUGH MAPS ---")
    
    // Create a map of student grades
    grades := map[string]int{
        "Alice": 85,
        "Bob":   92,
        "Carol": 78,
        "David": 88,
    }
    
    // Loop through map (order is random!)
    fmt.Println("Student Grades:")
    for name, grade := range grades {
        fmt.Printf("  %s: %d%%\n", name, grade)
    }
    
    // ==========================================
    // 6. LOOPING THROUGH STRINGS
    // ==========================================
    
    fmt.Println("\n--- 6. LOOPING THROUGH STRINGS ---")
    
    message := "Hello"
    
    // Loop through string characters
    fmt.Println("String characters:")
    for index, char := range message {
        fmt.Printf("  Index %d: Character '%c' (Unicode: %d)\n", index, char, char)
    }
    
    // ==========================================
    // 7. NESTED LOOPS
    // ==========================================
    
    fmt.Println("\n--- 7. NESTED LOOPS ---")
    
    // Create a multiplication table
    fmt.Println("Multiplication Table (1-3):")
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            result := i * j
            fmt.Printf("  %d × %d = %d\n", i, j, result)
        }
        fmt.Println() // Empty line between tables
    }
    
    // ==========================================
    // 8. LOOP CONTROL STATEMENTS
    // ==========================================
    
    fmt.Println("\n--- 8. LOOP CONTROL STATEMENTS ---")
    
    // BREAK example - exit loop early
    fmt.Println("Break example (stop at 5):")
    for i := 1; i <= 10; i++ {
        if i == 5 {
            fmt.Println("  Reached 5 - breaking!")
            break // Exit the entire loop
        }
        fmt.Printf("  Current: %d\n", i)
    }
    
    // CONTINUE example - skip current iteration
    fmt.Println("Continue example (skip even numbers):")
    for i := 1; i <= 5; i++ {
        if i % 2 == 0 { // If number is even
            continue // Skip to next iteration
        }
        fmt.Printf("  Odd number: %d\n", i)
    }
    
    // ==========================================
    // 9. PRACTICAL EXAMPLES
    // ==========================================
    
    fmt.Println("\n--- 9. PRACTICAL REAL-WORLD EXAMPLES ---")
    
    // Example 1: Sum of numbers
    numbers := []int{10, 20, 30, 40, 50}
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Printf("Sum of numbers %v = %d\n", numbers, sum)
    
    // Example 2: Find maximum value
    values := []int{45, 12, 78, 34, 89, 23}
    max := values[0] // Start with first element
    for _, value := range values {
        if value > max {
            max = value
        }
    }
    fmt.Printf("Maximum value in %v = %d\n", values, max)
    
    // Example 3: Filter even numbers
    allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var evenNumbers []int
    
    for _, num := range allNumbers {
        if num % 2 == 0 {
            evenNumbers = append(evenNumbers, num)
        }
    }
    fmt.Printf("Even numbers from %v = %v\n", allNumbers, evenNumbers)
    
    // Example 4: Count character frequency
    text := "programming"
    frequency := make(map[rune]int)
    
    for _, char := range text {
        frequency[char]++ // Count each character
    }
    
    fmt.Println("Character frequency in 'programming':")
    for char, count := range frequency {
        fmt.Printf("  '%c': %d times\n", char, count)
    }
    
    // ==========================================
    // 10. ADVANCED LOOP PATTERNS
    // ==========================================
    
    fmt.Println("\n--- 10. ADVANCED LOOP PATTERNS ---")
    
    // Pattern 1: Loop with multiple variables
    fmt.Println("Multiple variables in loop:")
    for i, j := 0, 10; i < 3 && j > 5; i, j = i+1, j-2 {
        fmt.Printf("  i=%d, j=%d\n", i, j)
    }
    
    // Pattern 2: Labeled break (break outer loop from inner loop)
    fmt.Println("Labeled break example:")
    OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                fmt.Println("  Breaking outer loop from inner loop")
                break OuterLoop // Break the outer loop, not just inner
            }
            fmt.Printf("  i=%d, j=%d\n", i, j)
        }
    }
    
    // Pattern 3: Loop with deferred cleanup
    fmt.Println("Loop with condition check:")
    for i := 0; i < 6; i++ {
        if i == 0 {
            continue // Skip first iteration
        }
        if i == 4 {
            break // Stop at 4
        }
        fmt.Printf("  Processing: %d\n", i)
    }
    
    // ==========================================
    // 11. COMMON PITFALLS AND SOLUTIONS
    // ==========================================
    
    fmt.Println("\n--- 11. COMMON PITFALLS ---")
    
    // Pitfall 1: Modifying slice while iterating
    numbersSlice := []int{1, 2, 3, 4, 5}
    fmt.Println("Original slice:", numbersSlice)
    
    // Safe way: Create a new slice or iterate backwards
    var filtered []int
    for _, num := range numbersSlice {
        if num != 3 { // Filter out number 3
            filtered = append(filtered, num)
        }
    }
    fmt.Println("After filtering out 3:", filtered)
    
    // Pitfall 2: Infinite loops
    /*
    // ❌ DANGEROUS - Infinite loop (commented out to prevent hanging)
    for {
        fmt.Println("This would run forever!")
    }
    */
    
    // ✅ Safe alternative with break condition
    safeCounter := 0
    for {
        safeCounter++
        if safeCounter >= 3 {
            break
        }
    }
    fmt.Println("Safe loop completed without hanging")
    
    fmt.Println("\n=== LOOP DEMONSTRATION COMPLETE ===")



//     Expected Output:

//     === LOOPS IN GOLANG DEMONSTRATION ===

// --- 1. BASIC FOR LOOP ---
// Basic loop - Iteration 0
// Basic loop - Iteration 1
// Basic loop - Iteration 2
// Basic loop - Iteration 3
// Basic loop - Iteration 4

// --- 2. WHILE-STYLE LOOP ---
// While-style - Count: 0
// While-style - Count: 1
// While-style - Count: 2

// --- 3. INFINITE LOOP WITH BREAK ---
// Infinite loop - Counter: 0
// Infinite loop - Counter: 1
// Infinite loop - Counter: 2
// Breaking out of infinite loop!

// ... and much more!
}