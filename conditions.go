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
    // CONDITIONAL STATEMENTS IN GOLANG
    // ==========================================

    fmt.Println("=== CONDITIONAL STATEMENTS DEMONSTRATION ===")

    // Example 1: Simple if statement with comparison
    x := 10
    y := 20

    // Compare if x is less than y
    val := x < y  // This evaluates to true (10 < 20)
    fmt.Println("Is x less than y?:", val)

    // ==========================================
    // BASIC IF STATEMENTS
    // ==========================================

    fmt.Println("\n--- BASIC IF STATEMENTS ---")

    // Example 2: Simple if condition
    age := 18
    if age >= 18 {
        fmt.Println("✅ You are an adult (age:", age, ")")
    }

    // Example 3: If-else statement
    temperature := 25
    if temperature > 30 {
        fmt.Println("🔥 It's hot outside!")
    } else {
        fmt.Println("😊 Weather is comfortable")
    }

    // Example 4: If-else if-else chain
    score := 85
    if score >= 90 {
        fmt.Println("🎉 Grade: A")
    } else if score >= 80 {
        fmt.Println("👍 Grade: B")
    } else if score >= 70 {
        fmt.Println("👌 Grade: C")
    } else {
        fmt.Println("💪 Grade: D - Needs improvement")
    }

    // ==========================================
    // COMPARISON OPERATORS
    // ==========================================

    fmt.Println("\n--- COMPARISON OPERATORS ---")

    a := 15
    b := 10

    // Demonstrate all comparison operators
    fmt.Printf("a = %d, b = %d\n", a, b)
    fmt.Println("a == b (equal):", a == b)        // false
    fmt.Println("a != b (not equal):", a != b)    // true  
    fmt.Println("a < b (less than):", a < b)      // false
    fmt.Println("a <= b (less or equal):", a <= b) // false
    fmt.Println("a > b (greater than):", a > b)   // true
    fmt.Println("a >= b (greater or equal):", a >= b) // true

    // ==========================================
    // LOGICAL OPERATORS (AND, OR, NOT)
    // ==========================================

    fmt.Println("\n--- LOGICAL OPERATORS ---")

    // Example 5: AND operator (&&) - both conditions must be true
    hasTicket := true
    hasID := true
    if hasTicket && hasID {
        fmt.Println("✅ Welcome to the event!")
    } else {
        fmt.Println("❌ Cannot enter the event")
    }

    // Example 6: OR operator (||) - at least one condition must be true
    isWeekend := false
    isHoliday := true
    if isWeekend || isHoliday {
        fmt.Println("🎉 No work today!")
    } else {
        fmt.Println("💼 Time to work")
    }

    // Example 7: NOT operator (!) - reverses the condition
    isRaining := false
    if !isRaining {
        fmt.Println("☀️ Great weather for a walk!")
    } else {
        fmt.Println("🌧️ Better stay indoors")
    }

    // Example 8: Complex logical conditions
    userAge := 25
    hasLicense := true
    hasCar := false

    if (userAge >= 18 && hasLicense) || hasCar {
        fmt.Println("🚗 You can drive")
    } else {
        fmt.Println("🚫 You cannot drive")
    }

    // ==========================================
    // VARIABLE SCOPE IN IF STATEMENTS
    // ==========================================

    fmt.Println("\n--- VARIABLE SCOPE IN IF ---")

    // Example 9: Variable declaration in if statement
    // The variable 'discount' only exists within the if block
    price := 100.0
    if discount := 0.1; price > 50 {  // discount is declared and initialized here
        finalPrice := price - (price * discount)
        fmt.Printf("Original: $%.2f, Discount: %.0f%%, Final: $%.2f\n", 
                   price, discount*100, finalPrice)
    }
    // discount variable is not accessible here - it's out of scope

    // Example 10: Multiple conditions in if
    number := 15
    if number > 10 && number < 20 && number % 5 == 0 {
        fmt.Printf("✅ %d is between 10-20 AND divisible by 5\n", number)
    }

    // ==========================================
    // NESTED IF STATEMENTS
    // ==========================================

    fmt.Println("\n--- NESTED IF STATEMENTS ---")

    // Example 11: Nested if statements
    username := "admin"
    password := "12345"

    if username == "admin" {
        if password == "12345" {
            fmt.Println("🔓 Login successful!")
        } else {
            fmt.Println("❌ Wrong password")
        }
    } else {
        fmt.Println("❌ User not found")
    }

    // ==========================================
    // REAL-WORLD PRACTICAL EXAMPLES
    // ==========================================

    fmt.Println("\n--- REAL-WORLD EXAMPLES ---")

    // Example 12: Grade calculator
    testScore := 78
    attendance := 85

    if testScore >= 60 && attendance >= 75 {
        fmt.Println("🎓 You passed the course!")
        if testScore >= 90 {
            fmt.Println("🏆 Excellent work!")
        } else if testScore >= 80 {
            fmt.Println("👍 Good job!")
        } else {
            fmt.Println("👏 You made it!")
        }
    } else {
        fmt.Println("📚 You need to improve")
        if testScore < 60 {
            fmt.Println("   - Study harder for tests")
        }
        if attendance < 75 {
            fmt.Println("   - Improve your attendance")
        }
    }

    // Example 13: Shopping cart discount
    cartTotal := 120.0
    isMember := true

    var discountRate float64
    if isMember {
        if cartTotal > 100 {
            discountRate = 0.2  // 20% discount for members over $100
        } else {
            discountRate = 0.1  // 10% discount for members
        }
    } else {
        if cartTotal > 100 {
            discountRate = 0.1  // 10% discount for non-members over $100
        } else {
            discountRate = 0.0  // No discount
        }
    }

    finalAmount := cartTotal - (cartTotal * discountRate)
    fmt.Printf("\n🛒 Shopping Cart Summary:\n")
    fmt.Printf("   Subtotal: $%.2f\n", cartTotal)
    fmt.Printf("   Member: %v\n", isMember)
    fmt.Printf("   Discount: %.0f%%\n", discountRate*100)
    fmt.Printf("   Total: $%.2f\n", finalAmount)

    // ==========================================
    // TERNARY OPERATOR ALTERNATIVE
    // ==========================================

    fmt.Println("\n--- TERNARY-STYLE CONDITIONS ---")

    // Go doesn't have ternary operator, but we can simulate it
    maxValue := 100
    userInput := 150

    // Using if-else to assign value (ternary alternative)
    result := userInput
    if userInput > maxValue {
        result = maxValue
    }
    fmt.Printf("User input: %d, Final value: %d\n", userInput, result)

    // Another way using immediate if evaluation
    finalScore := 65
    status := "Fail"
    if finalScore >= 50 {
        status = "Pass"
    }
    fmt.Printf("Score: %d, Status: %s\n", finalScore, status)

    // ==========================================
    // COMMON PITFALLS AND TIPS
    // ==========================================

    fmt.Println("\n--- COMMON PITFALLS ---")

    // Pitfall 1: Assignment (=) vs Comparison (==)
    var isActive bool = true
    
    // ❌ WRONG: if isActive = false { ... } (this assigns, not compares)
    // ✅ CORRECT: if isActive == false { ... } or if !isActive { ... }
    
    if !isActive {
        fmt.Println("Account is inactive")
    } else {
        fmt.Println("Account is active")
    }

    // Pitfall 2: Floating point comparisons
    float1 := 0.1 + 0.2
    float2 := 0.3
    
    // ❌ Don't do this: if float1 == float2
    // ✅ Do this instead (with tolerance):
    tolerance := 0.0001
    if (float1 - float2) < tolerance && (float2 - float1) < tolerance {
        fmt.Printf("✅ Floats are approximately equal: %.17f ≈ %.17f\n", float1, float2)
    } else {
        fmt.Printf("❌ Floats are not equal: %.17f ≠ %.17f\n", float1, float2)
    }

    fmt.Println("\n=== CONDITIONALS DEMONSTRATION COMPLETE ===")


//     Expected Output:

//     === CONDITIONAL STATEMENTS DEMONSTRATION ===
// Is x less than y?: true

// --- BASIC IF STATEMENTS ---
// ✅ You are an adult (age: 18 )
// 😊 Weather is comfortable
// 👍 Grade: B

// --- COMPARISON OPERATORS ---
// a = 15, b = 10
// a == b (equal): false
// a != b (not equal): true
// a < b (less than): false
// a <= b (less or equal): false
// a > b (greater than): true
// a >= b (greater or equal): true

// --- LOGICAL OPERATORS ---
// ✅ Welcome to the event!
// 🎉 No work today!
// ☀️ Great weather for a walk!
// 🚗 You can drive

// ... and more!
}