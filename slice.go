package main

import "fmt"

func main() {
    // ========== 1. BASIC SLICE CREATION ==========
    fmt.Println("=== 1. BASIC SLICE CREATION ===")
    
    // Method 1: From an array
    var x [6]int = [6]int{1, 2, 3, 4, 5, 6}
    var s []int = x[1:3]  // Creates slice from index 1 to 2 (exclusive of 3)
    fmt.Println("Slice s:", s)                    // [2 3]
    fmt.Println("Full capacity:", s[:cap(s)])     // [2 3 4 5 6]
    
    // Method 2: Using slice literal
    slice1 := []int{10, 20, 30, 40, 50}
    fmt.Println("Slice literal:", slice1)         // [10 20 30 40 50]
    
    // Method 3: Using make() - specify length and capacity
    slice2 := make([]int, 3, 5)  // length=3, capacity=5
    fmt.Printf("Make slice: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))
    
    // ========== 2. SLICE OPERATIONS ==========
    fmt.Println("\n=== 2. SLICE OPERATIONS ===")
    
    // Appending to slices (FIXING YOUR CODE)
    var a []string = []string{"yes", "no", "come"}
    a = append(a, "let")  // append returns new slice, must assign back
    a = append(a, "go", "stop")
    fmt.Println("After append:", a)  // [yes no come let go stop]
    
    // Slicing operations
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("Original:", numbers)
    fmt.Println("numbers[2:5]:", numbers[2:5])    // [3 4 5] - index 2 to 4
    fmt.Println("numbers[:4]:", numbers[:4])      // [1 2 3 4] - start to index 3
    fmt.Println("numbers[6:]:", numbers[6:])      // [7 8 9 10] - index 6 to end
    fmt.Println("numbers[:]:", numbers[:])        // [1 2 3 4 5 6 7 8 9 10] - full slice
    
    // ========== 3. LENGTH vs CAPACITY ==========
    fmt.Println("\n=== 3. LENGTH vs CAPACITY ===")
    
    arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    sliceFromArr := arr[2:5]  // elements at index 2,3,4
    
    fmt.Printf("Array: %v\n", arr)
    fmt.Printf("Slice: %v, len=%d, cap=%d\n", sliceFromArr, len(sliceFromArr), cap(sliceFromArr))
    // Slice: [2 3 4], len=3, cap=8 (can expand to use arr[2] to arr[9])
    
    // Showing full capacity
    fmt.Println("Using full capacity:", sliceFromArr[:cap(sliceFromArr)])
    // [2 3 4 5 6 7 8 9]
    
    // ========== 4. MODIFYING SLICES ==========
    fmt.Println("\n=== 4. MODIFYING SLICES ===")
    
    // Slices are references to underlying arrays
    original := []int{100, 200, 300}
    reference := original[1:]  // [200, 300]
    reference[0] = 999         // This modifies the underlying array!
    fmt.Println("Original after modification:", original)  // [100 999 300]
    
    // ========== 5. COPYING SLICES ==========
    fmt.Println("\n=== 5. COPYING SLICES ===")
    
    // Method 1: Using copy() function
    src := []int{1, 2, 3, 4, 5}
    dest := make([]int, len(src))
    copy(dest, src)
    dest[0] = 99
    fmt.Println("Source:", src)    // [1 2 3 4 5] - unchanged
    fmt.Println("Destination:", dest) // [99 2 3 4 5] - modified
    
    // Method 2: Using append trick
    copy2 := append([]int{}, src...)
    copy2[1] = 88
    fmt.Println("Source:", src)      // [1 2 3 4 5] - unchanged
    fmt.Println "Copied slice:", copy2) // [1 88 3 4 5]
    
    // ========== 6. SLICE GROWTH AND REALLOCATION ==========
    fmt.Println("\n=== 6. SLICE GROWTH ===")
    
    var growth []int
    fmt.Printf("Initial - len=%d, cap=%d\n", len(growth), cap(growth))
    
    for i := 0; i < 10; i++ {
        growth = append(growth, i)
        fmt.Printf("After appending %d - len=%d, cap=%d\n", i, len(growth), cap(growth))
    }
    // Notice how capacity doubles when needed: 0 → 1 → 2 → 4 → 8 → 16
    
    // ========== 7. MULTI-DIMENSIONAL SLICES ==========
    fmt.Println("\n=== 7. MULTI-DIMENSIONAL SLICES ===")
    
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("2D slice:")
    for i, row := range matrix {
        fmt.Printf("Row %d: %v\n", i, row)
    }
    
    // ========== 8. PRACTICAL EXAMPLES ==========
    fmt.Println("\n=== 8. PRACTICAL EXAMPLES ===")
    
    // Removing elements
    slice := []int{10, 20, 30, 40, 50}
    fmt.Println("Original:", slice)
    
    // Remove element at index 2
    slice = append(slice[:2], slice[3:]...)
    fmt.Println("After removing index 2:", slice) // [10 20 40 50]
    
    // Inserting elements
    slice = []int{1, 2, 3, 4, 5}
    // Insert 99 at index 2
    slice = append(slice[:2], append([]int{99}, slice[2:]...)...)
    fmt.Println("After inserting 99 at index 2:", slice) // [1 2 99 3 4 5]
    
    // ========== 9. COMMON PITFALLS ==========
    fmt.Println("\n=== 9. COMMON PITFALLS ===")
    
    // Pitfall 1: Forgetting to assign append result
    wrongSlice := []int{1, 2, 3}
    append(wrongSlice, 4)  // ❌ This does nothing!
    fmt.Println("Wrong way - slice unchanged:", wrongSlice)
    
    correctSlice := []int{1, 2, 3}
    correctSlice = append(correctSlice, 4)  // ✅ Correct!
    fmt.Println("Correct way - slice updated:", correctSlice)
    
    // Pitfall 2: Memory sharing
    base := []int{1, 2, 3, 4, 5}
    part1 := base[1:3]  // [2, 3]
    part2 := base[2:4]  // [3, 4]
    
    part1[1] = 999  // This affects both part1 and part2!
    fmt.Println("base:", base)   // [1 2 999 4 5]
    fmt.Println("part1:", part1) // [2 999]
    fmt.Println("part2:", part2) // [999 4]
}

// ========== 10. HELPER FUNCTIONS ==========
func printSliceInfo(name string, s []int) {
    fmt.Printf("%s: %v, len=%d, cap=%d\n", name, s, len(s), cap(s))
}