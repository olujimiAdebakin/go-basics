package main

import "fmt"

// Visual Representation:
// Memory:
// Address: 0xc000018030
// Value:   7
//          ↑
//          y points here

// x ←───┐
//       │ (both access same memory)
// *y ←──┘

func main() {

	x := 10        // normal variable
	ptr := &x      // &x means "address of x"

	fmt.Println("x:", x)
	fmt.Println("ptr:", ptr)      // prints memory address like 0xc000014090
	fmt.Println("value at ptr:", *ptr) // *ptr means "get the value at this address"


    fmt.Println("=== POINTERS AND DEREFERENCING ===")
    
    // ========== BASIC POINTER OPERATIONS ==========
    fmt.Println("\n1. BASIC POINTER OPERATIONS")
    
    x := 7
    fmt.Printf("x = %d (value)\n", x)
    fmt.Printf("&x = %p (memory address of x)\n", &x)
    
    // y is a pointer to x (stores the memory address of x)
    y := &x
    fmt.Printf("y = %p (memory address stored in y)\n", y)
    fmt.Printf("*y = %d (value at the address stored in y)\n", *y)
    
    // ========== MODIFYING THROUGH POINTERS ==========
    fmt.Println("\n2. MODIFYING THROUGH POINTERS")
    
    // Change x through the pointer y
    *y = 42  // This changes x to 42!
    fmt.Printf("After *y = 42:\n")
    fmt.Printf("x = %d, *y = %d\n", x, *y)
    
    // ========== POINTER COMPARISON ==========
    fmt.Println("\n3. POINTER COMPARISON")
    
    z := &x
    fmt.Printf("y = %p, z = %p\n", y, z)
    fmt.Printf("y == z? %t (both point to same address)\n", y == z)
    fmt.Printf("*y == *z? %t (both have same value)\n", *y == *z)
    
    // ========== NIL POINTERS ==========
    fmt.Println("\n4. NIL POINTERS")
    
    var ptr *int  // nil pointer (points to nothing)
    fmt.Printf("ptr = %p (nil pointer)\n", ptr)
    
    // This would cause runtime panic:
    // fmt.Printf("*ptr = %d\n", *ptr) // ❌ PANIC!
    
    // Safe nil check
    if ptr == nil {
        fmt.Println("ptr is nil - cannot dereference!")
    }
    
    // ========== POINTERS WITH STRUCTS ==========
    fmt.Println("\n5. POINTERS WITH STRUCTS")
    
    type Person struct {
        Name string
        Age  int
    }
    
    alice := Person{"Alice", 25}
    ptrToAlice := &alice
    
    fmt.Printf("alice: %+v\n", alice)
    fmt.Printf("ptrToAlice: %p\n", ptrToAlice)
    fmt.Printf("*ptrToAlice: %+v\n", *ptrToAlice)
    
    // Modify struct through pointer
    ptrToAlice.Age = 26  // Same as (*ptrToAlice).Age = 26
    fmt.Printf("After modification: %+v\n", alice)
    
    // ========== POINTERS IN FUNCTIONS ==========
    fmt.Println("\n6. POINTERS IN FUNCTIONS")
    
    a := 10
    b := 20
    fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
    swap(&a, &b)  // Pass addresses
    fmt.Printf("After swap: a=%d, b=%d\n", a, b)
    
    // ========== POINTER ARITHMETIC (NOT IN GO) ==========
    fmt.Println("\n7. NO POINTER ARITHMETIC IN GO")
    
    numbers := [3]int{1, 2, 3}
    ptr1 := &numbers[0]
    ptr2 := &numbers[1]
    
    fmt.Printf("ptr1 = %p (address of numbers[0])\n", ptr1)
    fmt.Printf("ptr2 = %p (address of numbers[1])\n", ptr2)
    fmt.Println("In Go, you CANNOT do: ptr1 + 1 (no pointer arithmetic)")
    
    // ========== PRACTICAL USE CASES ==========
    fmt.Println("\n8. PRACTICAL USE CASES")
    
    // Case 1: Large struct efficiency
    largeData := make([]int, 1000000) // Large slice
    processLargeData(&largeData)      // Pass pointer to avoid copying
    
    // Case 2: Modifying function arguments
    counter := 0
    increment(&counter)
    increment(&counter)
    fmt.Printf("Counter after increments: %d\n", counter)
    
    // Case 3: Optional parameters
    var optional *string
    printMessage(optional)                    // nil pointer
    msg := "Hello World"
    printMessage(&msg)                       // valid pointer
}

// ========== HELPER FUNCTIONS ==========

// Function that swaps values using pointers
func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
    fmt.Printf("Inside swap: *a=%d, *b=%d\n", *a, *b)
}

// Function that processes large data efficiently
func processLargeData(data *[]int) {
    // Work with the original data without copying
    (*data)[0] = 999
    fmt.Printf("Processed data - first element: %d\n", (*data)[0])
}

// Function that increments a counter
func increment(count *int) {
    *count += 1
}

// Function with optional parameter using pointer
func printMessage(msg *string) {
    if msg == nil {
        fmt.Println("No message provided")
    } else {
        fmt.Printf("Message: %s\n", *msg)
    }
}

// ========== COMMON PATTERNS ==========

func commonPatterns() {
    fmt.Println("\n9. COMMON POINTER PATTERNS")
    
    // Pattern 1: New() function returns pointer
    ptr := new(int)  // Allocates memory and returns pointer
    *ptr = 100
    fmt.Printf("new(int) pattern: %d\n", *ptr)
    
    // Pattern 2: Method receivers with pointers
    rect := Rectangle{width: 10, height: 5}
    rect.Scale(2)  // Modifies the original
    fmt.Printf("Scaled rectangle: %+v\n", rect)
    
    // Pattern 3: Returning pointers from functions
    person := createPerson("Bob", 30)
    fmt.Printf("Created person: %+v\n", *person)
}

type Rectangle struct {
    width, height int
}

// Pointer receiver method (can modify the struct)
func (r *Rectangle) Scale(factor int) {
    r.width *= factor
    r.height *= factor
}

// Function that returns a pointer
func createPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}