package main

import (
	_ "bufio"  // The underscore _ means we import but don't use this package
	 "fmt"     // This package is used for input/output operations
	_ "os"     // Not used in this program (ignored with _)
	_ "strconv" // Not used in this program (ignored with _)
	_ "text/template/parse"  // _ Blank identifier ignoring the imported package
)

func main() {
    fmt.Println("=== ARRAYS IN GOLANG DEMONSTRATION ===")
    
    // ==========================================
    // 1. ARRAY DECLARATION AND INITIALIZATION
    // ==========================================
    
    fmt.Println("\n--- 1. ARRAY DECLARATION METHODS ---")
    
    // Method 1: Explicit declaration with type and size
    var arr [5]int = [5]int{10, 20, 30, 40, 50}
    fmt.Println("Array 1 (explicit):", arr)
    
    // Method 2: Short declaration with values
    arrr := [5]int{3, 4, 5, 6, 7}
    fmt.Println("Array 2 (short):", arrr)
    
    // Method 3: Declaration with specific index initialization
    arr3 := [5]int{0: 100, 2: 300, 4: 500}
    fmt.Println("Array 3 (index init):", arr3) // [100 0 300 0 500]
    
    // Method 4: Let compiler count elements
    arr4 := [...]int{1, 2, 3, 4, 5, 6} // Size is 6 (determined by elements)
    fmt.Println("Array 4 (compiler count):", arr4)
    
    // Method 5: Zero-value initialization (all elements set to zero)
    var arr5 [3]int
    fmt.Println("Array 5 (zero values):", arr5) // [0 0 0]
    
    // ==========================================
    // 2. ACCESSING AND MODIFYING ARRAY ELEMENTS
    // ==========================================
    
    fmt.Println("\n--- 2. ACCESSING AND MODIFYING ELEMENTS ---")
    
    // Access individual elements (zero-indexed)
    fmt.Printf("arr[0] = %d\n", arr[0]) // First element
    fmt.Printf("arr[4] = %d\n", arr[4]) // Last element
    
    // Modify array elements
    arr[0] = 100
    arr[1] = 200
    fmt.Println("After modification:", arr)
    
    // ==========================================
    // 3. ARRAY LENGTH AND ITERATION
    // ==========================================
    
    fmt.Println("\n--- 3. ARRAY LENGTH AND ITERATION ---")
    
    // Get array length using len() function
    fmt.Printf("Length of arr: %d\n", len(arr))
    fmt.Printf("Length of arr4: %d\n", len(arr4))
    
    // Method 1: Traditional for loop
    sum := 0
    for i := 0; i < len(arrr); i++ {
        sum += arrr[i]
    }
    fmt.Printf("Sum using traditional loop: %d\n", sum)
    
    // Method 2: Range-based for loop (index and value)
    sum2 := 0
    for index, value := range arrr {
        fmt.Printf("Index %d: Value %d\n", index, value)
        sum2 += value
    }
    fmt.Printf("Sum using range loop: %d\n", sum2)
    
    // Method 3: Range-based for loop (value only)
    sum3 := 0
    for _, value := range arrr { // _ ignores the index
        sum3 += value
    }
    fmt.Printf("Sum using range (values only): %d\n", sum3)
    
    // ==========================================
    // 4. MULTI-DIMENSIONAL ARRAYS
    // ==========================================
    
    fmt.Println("\n--- 4. MULTI-DIMENSIONAL ARRAYS ---")
    
    // 2D Array (matrix)
    var matrix [3][3]int = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("2D Array (Matrix):")
    for i := 0; i < len(matrix); i++ {
        fmt.Println("  ", matrix[i])
    }
    
    // Access specific element in 2D array
    fmt.Printf("Element at [1][2]: %d\n", matrix[1][2]) // 6
    
    // 3D Array
    var cube [2][2][2]int = [2][2][2]int{
        {
            {1, 2},
            {3, 4},
        },
        {
            {5, 6},
            {7, 8},
        },
    }
    fmt.Println("3D Array (Cube):", cube)
    
    // ==========================================
    // 5. ARRAY OPERATIONS AND UTILITIES
    // ==========================================
    
    fmt.Println("\n--- 5. ARRAY OPERATIONS ---")
    
    // Finding maximum and minimum values
    numbers := [6]int{45, 12, 78, 34, 89, 23}
    max := numbers[0]
    min := numbers[0]
    
    for _, value := range numbers {
        if value > max {
            max = value
        }
        if value < min {
            min = value
        }
    }
    fmt.Printf("Array: %v\n", numbers)
    fmt.Printf("Maximum value: %d\n", max)
    fmt.Printf("Minimum value: %d\n", min)
    
    // Counting even and odd numbers
    evenCount := 0
    oddCount := 0
    for _, num := range numbers {
        if num % 2 == 0 {
            evenCount++
        } else {
            oddCount++
        }
    }
    fmt.Printf("Even numbers: %d, Odd numbers: %d\n", evenCount, oddCount)
    
    // ==========================================
    // 6. ARRAY COMPARISON
    // ==========================================
    
    fmt.Println("\n--- 6. ARRAY COMPARISON ---")
    
    // Arrays can be compared if they have the same type and size
    arrA := [3]int{1, 2, 3}
    arrB := [3]int{1, 2, 3}
    arrC := [3]int{4, 5, 6}
    
    fmt.Printf("arrA == arrB: %t\n", arrA == arrB) // true
    fmt.Printf("arrA == arrC: %t\n", arrA == arrC) // false
    fmt.Printf("arrA != arrC: %t\n", arrA != arrC) // true
    
    // ==========================================
    // 7. ARRAY AS FUNCTION PARAMETERS
    // ==========================================
    
    fmt.Println("\n--- 7. ARRAYS IN FUNCTIONS ---")
    
    // Arrays are passed by VALUE (copied) to functions
    testArray := [3]int{10, 20, 30}
    fmt.Println("Original array before function:", testArray)
    modifyArray(testArray) // This won't change the original
    fmt.Println("Original array after function:", testArray)
    
    // To modify original array, use pointers
    modifyArrayWithPointer(&testArray)
    fmt.Println("Original array after pointer function:", testArray)
    
    // ==========================================
    // 8. PRACTICAL REAL-WORLD EXAMPLES
    // ==========================================
    
    fmt.Println("\n--- 8. PRACTICAL EXAMPLES ---")
    
    // Example 1: Student Grades System
    var grades [5]int = [5]int{85, 92, 78, 96, 88}
    total := 0
    for _, grade := range grades {
        total += grade
    }
    average := float64(total) / float64(len(grades))
    fmt.Printf("Grades: %v\n", grades)
    fmt.Printf("Average grade: %.2f\n", average)
    
    // Example 2: Temperature Tracking
    var temperatures [7]float64 = [7]float64{22.5, 23.1, 21.8, 24.3, 25.0, 22.7, 23.9}
    fmt.Printf("Weekly temperatures: %.1f°C\n", temperatures)
    
    // Find hottest and coldest days
    hottest := temperatures[0]
    coldest := temperatures[0]
    for _, temp := range temperatures {
        if temp > hottest {
            hottest = temp
        }
        if temp < coldest {
            coldest = temp
        }
    }
    fmt.Printf("Hottest: %.1f°C, Coldest: %.1f°C\n", hottest, coldest)
    
    // Example 3: Inventory Management
    var inventory [4]string = [4]string{"Laptop", "Mouse", "Keyboard", "Monitor"}
    var prices [4]float64 = [4]float64{999.99, 25.50, 75.00, 299.99}
    
    fmt.Println("\nInventory List:")
    totalValue := 0.0
    for i := 0; i < len(inventory); i++ {
        fmt.Printf("  %s: $%.2f\n", inventory[i], prices[i])
        totalValue += prices[i]
    }
    fmt.Printf("Total inventory value: $%.2f\n", totalValue)
    
    // ==========================================
    // 9. ARRAY LIMITATIONS AND CHARACTERISTICS
    // ==========================================
    
    fmt.Println("\n--- 9. ARRAY CHARACTERISTICS ---")
    
    // Arrays have fixed size (cannot be resized)
    fixedArray := [3]int{1, 2, 3}
    fmt.Printf("Fixed array: %v (size: %d)\n", fixedArray, len(fixedArray))
    
    // Array type includes its size
    var size3 [3]int
    var size5 [5]int
    fmt.Printf("Type of size3: %T\n", size3) // [3]int
    fmt.Printf("Type of size5: %T\n", size5) // [5]int
    // size3 = size5 // ❌ ERROR: different types
    
    // ==========================================
    // 10. COPYING ARRAYS
    // ==========================================
    
    fmt.Println("\n--- 10. ARRAY COPYING ---")
    
    original := [3]int{10, 20, 30}
    copy := original // This creates a copy (not a reference)
    
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Copy: %v\n", copy)
    
    // Modify the copy
    copy[0] = 100
    fmt.Println("After modifying copy:")
    fmt.Printf("Original: %v (unchanged)\n", original)
    fmt.Printf("Copy: %v (changed)\n", copy)
    
    fmt.Println("\n=== ARRAY DEMONSTRATION COMPLETE ===")
}

// Function that receives array by value (copy)
func modifyArray(arr [3]int) {
    arr[0] = 999
    fmt.Println("Inside function (copy):", arr)
}

// Function that receives array by pointer (can modify original)
func modifyArrayWithPointer(arr *[3]int) {
    arr[0] = 999
    fmt.Println("Inside pointer function:", *arr)


//   Expected Output:  
//     === ARRAYS IN GOLANG DEMONSTRATION ===

// --- 1. ARRAY DECLARATION METHODS ---
// Array 1 (explicit): [10 20 30 40 50]
// Array 2 (short): [3 4 5 6 7]
// Array 3 (index init): [100 0 300 0 500]
// Array 4 (compiler count): [1 2 3 4 5 6]
// Array 5 (zero values): [0 0 0]

// --- 2. ACCESSING AND MODIFYING ELEMENTS ---
// arr[0] = 100
// arr[4] = 50
// After modification: [100 200 30 40 50]

// ... and much more!
}