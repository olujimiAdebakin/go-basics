package main

import (
	_ "bufio"  // The underscore _ means we import but don't use this package
	 "fmt"     // This package is used for input/output operations
	_ "os"     // Not used in this program (ignored with _)
	_ "strconv" // Not used in this program (ignored with _)
	_ "text/template/parse"  // _ Blank identifier ignoring the imported package
)

func main() {
    fmt.Println("=== SWITCH STATEMENTS IN GOLANG ===")
    
    // ==========================================
    // 1. BASIC SWITCH STATEMENT
    // ==========================================
    
    fmt.Println("\n--- 1. BASIC SWITCH STATEMENT ---")
    
    ans := 10

    // Basic switch with a variable
    switch ans {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 10:
        fmt.Println("ten") // This will execute since ans = 10
    default:
        fmt.Println("not a case")    
    }
    
    // ==========================================
    // 2. SWITCH WITH MULTIPLE VALUES PER CASE
    // ==========================================
    
    fmt.Println("\n--- 2. MULTIPLE VALUES PER CASE ---")
    
    day := "Monday"
    
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("🎉 It's the weekend!")
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("💼 It's a weekday")
    default:
        fmt.Println("❌ Invalid day")
    }
    
    // ==========================================
    // 3. SWITCH WITHOUT EXPRESSION (IF-ELSE ALTERNATIVE)
    // ==========================================
    
    fmt.Println("\n--- 3. SWITCH WITHOUT EXPRESSION ---")
    
    score := 85
    
    // Switch without expression acts like if-else chain
    switch {
    case score >= 90:
        fmt.Println("Grade: A 🏆")
    case score >= 80:
        fmt.Println("Grade: B 👍") // This will execute (85 >= 80)
    case score >= 70:
        fmt.Println("Grade: C 👌")
    case score >= 60:
        fmt.Println("Grade: D 📚")
    default:
        fmt.Println("Grade: F ❌")
    }
    
    // ==========================================
    // 4. SWITCH WITH FALLTHROUGH
    // ==========================================
    
    fmt.Println("\n--- 4. SWITCH WITH FALLTHROUGH ---")
    
    number := 2
    
    // fallthrough continues to the next case (even if condition is false)
    switch number {
    case 1:
        fmt.Println("Number is 1")
        fallthrough // Continue to next case
    case 2:
        fmt.Println("Number is 2") // This executes
        fallthrough // Continue to next case
    case 3:
        fmt.Println("Number is 3") // This also executes due to fallthrough
        // No fallthrough here, so it stops
    case 4:
        fmt.Println("Number is 4") // This won't execute
    }
    
    // ==========================================
    // 5. SWITCH WITH INITIALIZATION STATEMENT
    // ==========================================
    
    fmt.Println("\n--- 5. SWITCH WITH INITIALIZATION ---")
    
    // You can initialize a variable right in the switch statement
    switch hour := 14; { // Initialize hour and then switch on true
    case hour < 12:
        fmt.Println("Good morning! ☀️")
    case hour < 18:
        fmt.Println("Good afternoon! 🌤️") // This executes (14 < 18)
    default:
        fmt.Println("Good evening! 🌙")
    }
    // hour variable is not accessible here - it's scoped to the switch
    
    // ==========================================
    // 6. TYPE SWITCH (CHECKING VARIABLE TYPE)
    // ==========================================
    
    fmt.Println("\n--- 6. TYPE SWITCH ---")
    
    // Function to check type of any variable
    checkType := func(x interface{}) {
        switch x.(type) { // Special syntax for type checking
        case int:
            fmt.Println("🔢 Integer type")
        case string:
            fmt.Println("🔤 String type")
        case float64:
            fmt.Println("📊 Float type")
        case bool:
            fmt.Println("✅ Boolean type")
        default:
            fmt.Println("❓ Unknown type")
        }
    }
    
    checkType(42)        // Integer
    checkType("hello")   // String
    checkType(3.14)      // Float
    checkType(true)      // Boolean
    
    // ==========================================
    // 7. SWITCH WITH BREAK (EARLY EXIT)
    // ==========================================
    
    fmt.Println("\n--- 7. SWITCH WITH BREAK ---")
    
    value := 5
    
    switch value {
    case 1, 3, 5, 7, 9:
        fmt.Println("Odd number detected")
        if value == 5 {
            fmt.Println("Special case for 5 - breaking early!")
            break // Exit the switch immediately
        }
        fmt.Println("This won't print for value 5 due to break")
    case 2, 4, 6, 8, 10:
        fmt.Println("Even number detected")
    }
    
    // ==========================================
    // 8. PRACTICAL REAL-WORLD EXAMPLES
    // ==========================================
    
    fmt.Println("\n--- 8. PRACTICAL EXAMPLES ---")
    
    // Example 1: Calculator
    a, b := 15, 3
    operator := "/"
    
    switch operator {
    case "+":
        fmt.Printf("%d + %d = %d\n", a, b, a+b)
    case "-":
        fmt.Printf("%d - %d = %d\n", a, b, a-b)
    case "*":
        fmt.Printf("%d * %d = %d\n", a, b, a*b)
    case "/":
        if b != 0 {
            fmt.Printf("%d / %d = %d\n", a, b, a/b)
        } else {
            fmt.Println("Error: Division by zero!")
        }
    default:
        fmt.Println("Unknown operator")
    }
    
    // Example 2: Traffic Light System
    lightColor := "yellow"
    
    switch lightColor {
    case "red":
        fmt.Println("🚨 STOP!")
    case "yellow":
        fmt.Println("⚠️  SLOW DOWN!") // This executes
    case "green":
        fmt.Println("✅ GO!")
    default:
        fmt.Println("🚦 MALFUNCTION")
    }
    
    // Example 3: User Role Permissions
    userRole := "admin"
    
    switch userRole {
    case "admin":
        fmt.Println("🔐 Full access granted")
        fallthrough // Admin also gets moderator privileges
    case "moderator":
        fmt.Println("👮 Content moderation rights")
        fallthrough // Moderator also gets user privileges
    case "user":
        fmt.Println("👤 Basic user access")
    case "guest":
        fmt.Println("🚶 Read-only access")
    }
    
    // ==========================================
    // 9. COMPLEX CONDITIONS IN SWITCH
    // ==========================================
    
    fmt.Println("\n--- 9. COMPLEX CONDITIONS ---")
    
    temperature := 25
    humidity := 60
    
    // Switch with complex conditions
    switch {
    case temperature > 30 && humidity > 70:
        fmt.Println("🔥 Hot and humid - stay indoors!")
    case temperature > 30:
        fmt.Println("☀️ Hot but dry - stay hydrated")
    case temperature < 10:
        fmt.Println("❄️  Cold - wear warm clothes")
    case temperature >= 20 && temperature <= 30 && humidity < 50:
        fmt.Println("😊 Perfect weather!") // This executes
    default:
        fmt.Println("🤔 Average weather conditions")
    }
    
    // ==========================================
    // 10. SWITCH WITH FUNCTIONS
    // ==========================================
    
    fmt.Println("\n--- 10. SWITCH WITH FUNCTION CALLS ---")
    
    // You can call functions in switch cases
    getSeason := func(month int) string {
        switch month {
        case 12, 1, 2:
            return "Winter ❄️"
        case 3, 4, 5:
            return "Spring 🌸"
        case 6, 7, 8:
            return "Summer ☀️"
        case 9, 10, 11:
            return "Autumn 🍂"
        default:
            return "Invalid month"
        }
    }
    
    fmt.Println("Month 3 is in:", getSeason(3))
    fmt.Println("Month 8 is in:", getSeason(8))
    
    // ==========================================
    // 11. COMMON PITFALLS AND TIPS
    // ==========================================
    
    fmt.Println("\n--- 11. COMMON PITFALLS ---")
    
    // Tip 1: Switch cases don't need break (unlike other languages)
    // In Go, switch cases automatically break
    
    color := "red"
    switch color {
    case "red":
        fmt.Println("Color is red") // This executes
        // No need for break - it won't fall through to next case
    case "blue":
        fmt.Println("Color is blue") // This won't execute
    }
    
    // Tip 2: Use fallthrough only when explicitly needed
    fmt.Println("\nFallthrough example:")
    x := 1
    switch x {
    case 1:
        fmt.Println("Case 1")
        fallthrough // Explicit fallthrough
    case 2:
        fmt.Println("Case 2") // This also executes due to fallthrough
    }
    
    fmt.Println("\n=== SWITCH STATEMENT DEMONSTRATION COMPLETE ===")


//     Expected Output:

//     === SWITCH STATEMENTS IN GOLANG ===

// --- 1. BASIC SWITCH STATEMENT ---
// ten

// --- 2. MULTIPLE VALUES PER CASE ---
// 💼 It's a weekday

// --- 3. SWITCH WITHOUT EXPRESSION ---
// Grade: B 👍

// --- 4. SWITCH WITH FALLTHROUGH ---
// Number is 2
// Number is 3

// ... and much more!
}