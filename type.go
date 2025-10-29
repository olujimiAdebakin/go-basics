package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	 _ "text/template/parse"  // ✅ Blank identifier ignoring the imported package
)


func main() {
	
	fmt.Println("Program starting...")
	type_convert() // Call your function
	fmt.Println("Program completed!")
}



func type_convert() {

	scanner := bufio.NewScanner(os.Stdin)

	// First example: Read a line of input and print it back
	fmt.Printf("Please enter some input: %s", "\n")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You entered: %q\n", input)


	
	// Second example: Read year of birth and calculate age in 2025
	fmt.Printf("Type your DOB: %s", "\n")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2025:", 2025 - age)
}