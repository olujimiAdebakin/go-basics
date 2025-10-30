package main

import ( 
	"fmt" 
	"time"
)


// defer is a special keyword in Go that postpones the execution of a function until the surrounding function completes. Think of it like a "do this later" instruction!



func main() {
    fmt.Println("=== BASIC DEFER EXAMPLES ===")
    
    // Example 1: Simple defer
    fmt.Println("🎯 Example 1: Simple Defer")
    defer fmt.Println("Goodbye! 👋")  // This runs when main() ends
    fmt.Println("Hello! 🎉")
    
    // Example 2: Multiple defers (LIFO - Last In, First Out)
    fmt.Println("\n🎯 Example 2: Multiple Defers")
    defer fmt.Println("Defer 1 - Runs THIRD")
    defer fmt.Println("Defer 2 - Runs SECOND") 
    defer fmt.Println("Defer 3 - Runs FIRST")
    fmt.Println("Regular statement - Runs NOW")
    
    // Example 3: Defer in loops
    fmt.Println("\n🎯 Example 3: Defer in Loop")
    for i := 1; i <= 3; i++ {
        defer fmt.Printf("Defer in loop: %d\n", i)
    }
    fmt.Println("Loop finished")


     fmt.Println("\n=== PRACTICAL DEFER USE CASES ===")
    
    // Use Case 1: File operations (cleanup)
    fmt.Println("📁 Use Case 1: File Operations")
    openFile()
    
    // Use Case 2: Database connections (close connection)
    fmt.Println("\n🗄️ Use Case 2: Database Operations")
    databaseOperation()
    
    // Use Case 3: Unlocking mutexes
    fmt.Println("\n🔒 Use Case 3: Resource Locking")
    resourceLocking()
    
    // Use Case 4: Function timing
    fmt.Println("\n⏱️ Use Case 4: Function Timing")
    timedFunction()


// OUTPUT

// === BASIC DEFER EXAMPLES ===
// 🎯 Example 1: Simple Defer
// Hello! 🎉
// Goodbye! 👋

// 🎯 Example 2: Multiple Defers
// Regular statement - Runs NOW
// Defer 3 - Runs FIRST
// Defer 2 - Runs SECOND
// Defer 1 - Runs THIRD

// 🎯 Example 3: Defer in Loop
// Loop finished
// Defer in loop: 3
// Defer in loop: 2
// Defer in loop: 1

// === PRACTICAL DEFER USE CASES ===
// 📁 Use Case 1: File Operations
// Opening file...
// Reading file content...
// Processing data...
// Closing file...

// 🗄️ Use Case 2: Database Operations
// Connecting to database...
// Executing query...
// Fetching results...
// Disconnecting from database...

// 🔒 Use Case 3: Resource Locking
// Locking resource...
// Using the resource...
// Unlocking resource...

// ⏱️ Use Case 4: Function Timing
// Doing some work...
// Function took: 100.123ms
}




// Use Case 1: File operations
func openFile() {
    fmt.Println("Opening file...")
    defer fmt.Println("Closing file...")  // Always runs, even if error occurs
    
    fmt.Println("Reading file content...")
    fmt.Println("Processing data...")
    // File automatically "closed" when function ends
}

// Use Case 2: Database operations  
func databaseOperation() {
    fmt.Println("Connecting to database...")
    defer fmt.Println("Disconnecting from database...")
    
    fmt.Println("Executing query...")
    fmt.Println("Fetching results...")
    // If panic happens here, disconnect still runs!
}

// Use Case 3: Resource locking
func resourceLocking() {
    fmt.Println("Locking resource...")
    defer fmt.Println("Unlocking resource...")  // Always unlock!
    
    fmt.Println("Using the resource...")
    // Resource automatically unlocked when done
}

// Use Case 4: Function timing
func timedFunction() {
    start := time.Now()
    defer func() {
        duration := time.Since(start)
        fmt.Printf("Function took: %v\n", duration)
    }()
    
    fmt.Println("Doing some work...")
    time.Sleep(100 * time.Millisecond) // Simulate work
}