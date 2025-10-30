package main

import "fmt"

/*
MAKE EXPLANATION:
================

make() is a built-in Go function that ALLOCATES and INITIALIZES slices, maps, and channels.

Think of make() as a FACTORY that creates empty containers for your data:

🍬 make([]type, length, capacity) - Creates a SLICE (list of items)
🗺️ make(map[key]value)           - Creates a MAP (key-value storage)  
📞 make(chan type, buffer)       - Creates a CHANNEL (communication pipe)

Why use make()?
- Slices: Pre-allocate memory for better performance
- Maps: Must be initialized before use
- Channels: Required for goroutine communication

Syntax: make(Type, size...)
*/

func main() {
    fmt.Println("=== MAKE() COMPLETE GUIDE ===")
    
    // ========== 1. SLICE EXAMPLES ==========
    fmt.Println("\n1. SLICE EXAMPLES WITH make()")
    
    // Example 1: Guest list with pre-allocated capacity
    fmt.Println("🎯 Example 1: Guest List")
    guestList1 := make([]string, 0, 20)  // Empty slice that can hold 20 items
    fmt.Printf("With make: %v, len=%d, cap=%d\n", 
        guestList1, len(guestList1), cap(guestList1))
    
    // Example 2: Pre-filled slice with make
    fmt.Println("\n🎯 Example 2: Pre-filled Student List")
    students := make([]string, 3, 10)  // 3 empty spots, capacity for 10
    students[0] = "Alice"
    students[1] = "Bob"
    students[2] = "Charlie"
    fmt.Printf("Students: %v, len=%d, cap=%d\n", students, len(students), cap(students))
    
    // Example 3: Number sequence
    fmt.Println("\n🎯 Example 3: Number Sequence")
    numbers := make([]int, 5, 10)  // 5 zeros, capacity for 10
    for i := 0; i < 5; i++ {
        numbers[i] = i * 10
    }
    fmt.Printf("Numbers: %v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))
    
    // ========== 2. PERFORMANCE COMPARISON ==========
    fmt.Println("\n2. PERFORMANCE COMPARISON")
    
    // 🐢 Without pre-allocation (slower)
    fmt.Println("🐢 Without pre-allocation:")
    var slowList []string
    for i := 0; i < 10; i++ {
        slowList = append(slowList, fmt.Sprintf("Guest%d", i))
        fmt.Printf("  After %d appends: len=%d, cap=%d\n", i+1, len(slowList), cap(slowList))
    }
    
    // 🚀 With pre-allocation (faster)
    fmt.Println("\n🚀 With pre-allocation:")
    fastList := make([]string, 0, 10)
    for i := 0; i < 10; i++ {
        fastList = append(fastList, fmt.Sprintf("Guest%d", i))
        fmt.Printf("  After %d appends: len=%d, cap=%d\n", i+1, len(fastList), cap(fastList))
    }
    
    // ========== 3. MAP EXAMPLES ==========
    fmt.Println("\n3. MAP EXAMPLES WITH make()")
    
    // Example 4: Student grades
    fmt.Println("🎯 Example 4: Student Grade Book")
    gradeBook := make(map[string]int)  // Empty map for string keys and int values
    gradeBook["Alice"] = 95
    gradeBook["Bob"] = 87
    gradeBook["Charlie"] = 92
    fmt.Println("Grade Book:", gradeBook)
    
    // Example 5: Phone directory
    fmt.Println("\n🎯 Example 5: Phone Directory")
    phoneBook := make(map[string]string)
    phoneBook["Alice"] = "555-1234"
    phoneBook["Emergency"] = "911"
    phoneBook["Pizza"] = "555-1111"
    fmt.Println("Phone Book:", phoneBook)
    
    // Example 6: Shopping cart
    fmt.Println("\n🎯 Example 6: Shopping Cart")
    cart := make(map[string]float64)
    cart["Apple"] = 1.25
    cart["Bread"] = 3.50
    cart["Milk"] = 2.99
    fmt.Println("Shopping Cart:", cart)
    
    // ========== 4. CHANNEL EXAMPLES ==========
    fmt.Println("\n4. CHANNEL EXAMPLES WITH make()")
    
    // Example 7: Simple message channel
    fmt.Println("🎯 Example 7: Message Channel")
    messages := make(chan string)  // Unbuffered channel
    
    go func() {
        messages <- "Hello from goroutine!"
    }()
    
    msg := <-messages
    fmt.Println("Received:", msg)
    
    // Example 8: Buffered channel (mailbox)
    fmt.Println("\n🎯 Example 8: Buffered Channel (Mailbox)")
    mailbox := make(chan string, 3)  // Can hold 3 messages
    
    // Fill the mailbox
    mailbox <- "Letter 1"
    mailbox <- "Letter 2"
    mailbox <- "Letter 3"
    
    // Read from mailbox
    fmt.Println("Mailbox contains:")
    fmt.Println("  ", <-mailbox)
    fmt.Println("  ", <-mailbox)
    fmt.Println("  ", <-mailbox)
    
    // ========== 5. COMPARISON WITH OTHER METHODS ==========
    fmt.Println("\n5. COMPARISON WITH OTHER CREATION METHODS")
    
    // Method 1: make (you control capacity)
    guestList1 := make([]string, 0, 20)
    fmt.Printf("With make: %v, len=%d, cap=%d\n", 
        guestList1, len(guestList1), cap(guestList1))
    
    // Method 2: Slice literal (Go decides capacity)
    guestList2 := []string{}
    fmt.Printf("With literal: %v, len=%d, cap=%d\n", 
        guestList2, len(guestList2), cap(guestList2))
    
    // Method 3: var declaration (nil slice)
    var guestList3 []string
    fmt.Printf("With var: %v, len=%d, cap=%d\n", 
        guestList3, len(guestList3), cap(guestList3))
    
    // ========== 6. PRACTICAL USE CASES ==========
    fmt.Println("\n6. PRACTICAL USE CASES")
    
    // Use Case 1: Database results
    fmt.Println("📊 Use Case 1: Database Results")
    dbResults := make([]map[string]interface{}, 0, 100)  // Pre-allocate for 100 records
    // Simulate adding records
    for i := 0; i < 5; i++ {
        record := make(map[string]interface{})
        record["id"] = i
        record["name"] = fmt.Sprintf("User%d", i)
        dbResults = append(dbResults, record)
    }
    fmt.Println("Database results capacity:", cap(dbResults))
    fmt.Println("Database results:", dbResults)
    
    // Use Case 2: Configuration settings
    fmt.Println("\n⚙️ Use Case 2: Configuration Settings")
    config := make(map[string]interface{})
    config["app_name"] = "MyApp"
    config["version"] = 1.0
    config["debug"] = true
    config["max_users"] = 1000
    fmt.Println("Config:", config)
    
    // Use Case 3: Worker pool
    fmt.Println("\n👷 Use Case 3: Worker Pool Channels")
    jobs := make(chan int, 10)      // Buffered channel for jobs
    results := make(chan int, 10)   // Buffered channel for results
    
    // Start worker
    go func() {
        for job := range jobs {
            results <- job * 2  // Double the number
        }
    }()
    
    // Send jobs
    for i := 1; i <= 3; i++ {
        jobs <- i
    }
    close(jobs)
    
    // Get results
    fmt.Println("Worker results:")
    for i := 1; i <= 3; i++ {
        fmt.Printf("  Job %d → Result %d\n", i, <-results)
    }
    
    // ========== 7. COMMON MISTAKES ==========
    fmt.Println("\n7. COMMON MISTAKES")
    
    // ❌ Wrong: Forgetting make for maps
    var badMap map[string]int
    // badMap["key"] = 1  // This would PANIC at runtime!
    fmt.Println("❌ Bad map (nil):", badMap)
    
    // ✅ Correct: Using make for maps
    goodMap := make(map[string]int)
    goodMap["key"] = 1
    fmt.Println("✅ Good map:", goodMap)
    
    fmt.Println("\n=== END OF MAKE() GUIDE ===")
}

/*
KEY TAKEAWAYS:
=============
1. Use make() for SLICES when you know the capacity in advance
2. Use make() for MAPS always (they must be initialized)
3. Use make() for CHANNELS always (required for communication)
4. make() returns the actual type (not a pointer)
5. Pre-allocating with make() can significantly improve performance
*/