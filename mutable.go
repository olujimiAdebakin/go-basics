package main

import "fmt"


func main() {
    fmt.Println("\n=== MUTABLE EXAMPLES ===")
    
    // Example 1: Slices are mutable
    fmt.Println("🍬 SLICES ARE MUTABLE")
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Printf("Original: %v (memory: %p)\n", fruits, &fruits[0])
    
    // Modify the slice - changes the ORIGINAL
    fruits[0] = "orange"                    // Direct modification
    fruits = append(fruits, "grape")        // Can grow
    fmt.Printf("Modified: %v (memory: %p)\n", fruits, &fruits[0])
    
    // Example 2: Maps are mutable
    fmt.Println("\n🗺️ MAPS ARE MUTABLE")
    studentGrades := map[string]int{
        "Alice": 95,
        "Bob":   87,
    }
    fmt.Printf("Original: %v\n", studentGrades)
    
    // Modify the map - changes the ORIGINAL
    studentGrades["Alice"] = 98              // Update existing
    studentGrades["Charlie"] = 92            // Add new
    delete(studentGrades, "Bob")             // Remove
    fmt.Printf("Modified: %v\n", studentGrades)
    
    // Example 3: Arrays (careful - they behave differently!)
    fmt.Println("\n📦 ARRAYS - VALUE TYPES")
    arr1 := [3]int{1, 2, 3}
    arr2 := arr1  // Creates COPY of entire array
    
    arr2[0] = 99  // Only changes arr2
    fmt.Printf("arr1: %v, arr2: %v\n", arr1, arr2)
}