package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	_ "text/template/parse"  // ✅ Blank identifier ignoring the imported package
)

func main() {
	// This is where the program starts execution
	fmt.Println("🚀 Program starting...")
	type_convert() // Calling our main function that handles user input
	fmt.Println("✅ Program completed!")
}

func type_convert() {
	// Create a scanner to read input from keyboard
	scanner := bufio.NewScanner(os.Stdin)

	// ==========================================
	// EXAMPLE 1: Basic String Input
	// ==========================================
	fmt.Print("📝 Please enter your name: ") // Using Print instead of Printf for cleaner prompt
	scanner.Scan()                           // Wait for user to type and press Enter
	name := scanner.Text()                   // Get what the user typed as a string
	fmt.Printf("👋 Hello, %s!\n", name)      // Print greeting with their name

	// ==========================================
	// EXAMPLE 2: Integer Input with Age Calculation
	// ==========================================
	fmt.Print("🎂 Type your birth year (e.g., 1990): ")
	scanner.Scan()
	birthYear, err := strconv.ParseInt(scanner.Text(), 10, 64) // Convert string to number
	
	// Check if conversion worked
	if err != nil {
		fmt.Println("❌ Error: Please enter a valid year (numbers only!)")
	} else {
		age := 2025 - birthYear // Calculate age
		fmt.Printf("🎉 You will be %d years old at the end of 2025\n", age)
	}

	// ==========================================
	// EXAMPLE 3: Floating Point Number Input
	// ==========================================
	fmt.Print("💰 Enter your height in meters (e.g., 1.75): ")
	scanner.Scan()
	height, err := strconv.ParseFloat(scanner.Text(), 64) // Convert to decimal number
	if err != nil {
		fmt.Println("❌ Error: Please enter a valid height (e.g., 1.75)")
	} else {
		fmt.Printf("📏 Your height is %.2f meters\n", height) // %.2f shows 2 decimal places
	}

	// ==========================================
	// EXAMPLE 4: Boolean (Yes/No) Input
	// ==========================================
	fmt.Print("🎓 Are you a student? (yes/no): ")
	scanner.Scan()
	answer := scanner.Text()
	
	// Check the answer
	if answer == "yes" || answer == "Yes" || answer == "YES" {
		fmt.Println("📚 Great! Keep learning!")
	} else if answer == "no" || answer == "No" || answer == "NO" {
		fmt.Println("💼 Hope you're enjoying your work!")
	} else {
		fmt.Println("🤔 I didn't understand that answer")
	}

	// ==========================================
	// EXAMPLE 5: Multiple Inputs on One Line
	// ==========================================
	fmt.Print("🏠 Enter your city and country (e.g., London UK): ")
	scanner.Scan()
	var city, country string
	// Scanf can read multiple values at once
	fmt.Sscanf(scanner.Text(), "%s %s", &city, &country)
	fmt.Printf("🌍 You live in %s, %s\n", city, country)

	// ==========================================
	// EXAMPLE 6: Character Input (First Letter)
	// ==========================================
	fmt.Print("🔤 Enter your favorite letter: ")
	scanner.Scan()
	favoriteLetter := scanner.Text()
	if len(favoriteLetter) > 0 {
		firstChar := favoriteLetter[0] // Get first character
		fmt.Printf("⭐ Your favorite letter is '%c'\n", firstChar)
	}

	// ==========================================
	// EXAMPLE 7: Number Range Validation
	// ==========================================
	fmt.Print("🎯 Enter a number between 1 and 100: ")
	scanner.Scan()
	number, err := strconv.Atoi(scanner.Text()) // Convert to integer
	if err != nil {
		fmt.Println("❌ Please enter a valid number!")
	} else if number >= 1 && number <= 100 {
		fmt.Printf("✅ Great! %d is within range!\n", number)
	} else {
		fmt.Printf("❌ %d is not between 1 and 100\n", number)
	}

	// ==========================================
	// EXAMPLE 8: Password-style Input (Hidden)
	// ==========================================
	fmt.Print("🔒 Enter a secret word (it will show as stars): ")
	scanner.Scan()
	secret := scanner.Text()
	fmt.Printf("🔑 You entered: %s\n", "*****") // Don't show the actual secret
	fmt.Printf("📏 Secret length: %d characters\n", len(secret))

	// ==========================================
	// EXAMPLE 9: Simple Math Operation
	// ==========================================
	fmt.Print("➗ Enter two numbers to multiply (e.g., 5 3): ")
	scanner.Scan()
	var num1, num2 float64
	fmt.Sscanf(scanner.Text(), "%f %f", &num1, &num2)
	result := num1 * num2
	fmt.Printf("🧮 %.2f × %.2f = %.2f\n", num1, num2, result)

	// ==========================================
	// EXAMPLE 10: String Manipulation
	// ==========================================
	fmt.Print("✏️ Enter a sentence: ")
	scanner.Scan()
	sentence := scanner.Text()
	fmt.Printf("📊 Analysis:\n")
	fmt.Printf("   Original: %s\n", sentence)
	fmt.Printf("   UPPERCASE: %s\n", toUpper(sentence))
	fmt.Printf("   Length: %d characters\n", len(sentence))
	fmt.Printf("   Words: approximately %d words\n", len(sentence)/6) // Rough estimate
}

// Helper function to convert string to uppercase
func toUpper(s string) string {
	var result string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32) // Convert lowercase to uppercase
		} else {
			result += string(char) // Keep other characters as they are
		}
	}
	return result
	
	// Sample Output:

// 🚀 Program starting...
// 📝 Please enter your name: John
// 👋 Hello, John!
// 🎂 Type your birth year (e.g., 1990): 1995
// 🎉 You will be 30 years old at the end of 2025
// 💰 Enter your height in meters (e.g., 1.75): 1.80
// 📏 Your height is 1.80 meters
// ... and so on for other inputs

}

