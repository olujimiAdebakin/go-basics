package main

import "fmt"

func main() {
    fmt.Println("=== IMMUTABLE EXAMPLES ===")
    
    // Example 1: Strings are immutable
    fmt.Println("📝 STRINGS ARE IMMUTABLE")
    original := "Hello"
    fmt.Printf("Original: %s (memory: %p)\n", original, &original)
    
    modified := original + " World"  // Creates NEW string
    fmt.Printf("Modified: %s (memory: %p)\n", modified, &modified)
    
    // Trying to change a character directly:
    // original[0] = 'h'  // ❌ COMPILE ERROR! Strings are immutable
    
    // Example 2: Numbers are immutable
    fmt.Println("\n🔢 NUMBERS ARE IMMUTABLE")
    x := 10
    fmt.Printf("x = %d (memory: %p)\n", x, &x)
    
    y := x  // Creates COPY of the value
    y = 20  // Only changes y, not x
    fmt.Printf("x = %d, y = %d\n", x, y)
    
    // Example 3: Booleans are immutable
    fmt.Println("\n✅ BOOLEANS ARE IMMUTABLE")
    isActive := true
    wasActive := isActive  // Copy value
    isActive = false       // Only changes isActive
    
    fmt.Printf("isActive: %t, wasActive: %t\n", isActive, wasActive)
}