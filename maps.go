package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== COMPLETE MAP GUIDE ===")
	
	// ========== 1. BASIC MAP CREATION ==========
	fmt.Println("\n1. BASIC MAP CREATION")
	
	// Method 1: Declaration then initialization
	var studentAges map[string]int
	studentAges = make(map[string]int)
	studentAges["Alice"] = 23
	studentAges["Bob"] = 25
	
	// Method 2: Declare and initialize in one line
	capitals := make(map[string]string)
	capitals["USA"] = "Washington D.C."
	capitals["France"] = "Paris"
	capitals["Japan"] = "Tokyo"
	
	// Method 3: Using map literal (most common)
	var mp map[string]int = map[string]int{
		"apples":  1,
		"pear":    6,
		"orange":  8,
	}
	
	// Method 4: Short declaration with literal
	ap := map[int]uint{
		10: 20,
		20: 30,
		30: 40,
	}
	
	fmt.Println("Student Ages:", studentAges)
	fmt.Println("Capitals:", capitals)
	fmt.Println("Fruits:", mp)
	fmt.Println("Numbers:", ap)
	
	// ========== 2. ACCESSING AND MODIFYING ==========
	fmt.Println("\n2. ACCESSING AND MODIFYING MAPS")
	
	// Accessing values
	fmt.Println("Capital of France:", capitals["France"])
	fmt.Println("Apples count:", mp["apples"])
	
	// Modifying values
	mp["apples"] = 10
	ap[20] = 70
	fmt.Println("Updated fruits:", mp)
	fmt.Println("Updated numbers:", ap)
	
	// Adding new key-value pairs
	mp["banana"] = 5
	studentAges["Charlie"] = 22
	fmt.Println("After adding new items:")
	fmt.Println("Fruits:", mp)
	fmt.Println("Student Ages:", studentAges)
	
	// ========== 3. CHECKING KEY EXISTENCE ==========
	fmt.Println("\n3. CHECKING KEY EXISTENCE")
	
	// Two-value assignment to check existence
	val, ok := mp["apples"]
	fmt.Printf("mp['apples'] - Value: %d, Exists: %t\n", val, ok)
	
	val, ok = mp["watermelon"]
	fmt.Printf("mp['watermelon'] - Value: %d, Exists: %t\n", val, ok)
	
	// Shorter version in if statement
	if capital, exists := capitals["Germany"]; exists {
		fmt.Println("Capital of Germany:", capital)
	} else {
		fmt.Println("Germany not found in capitals map")
	}
	
	if age, ok := studentAges["Diana"]; ok {
		fmt.Println("Diana's age:", age)
	} else {
		fmt.Println("Diana not found in student ages")
	}
	
	// ========== 4. DELETING ITEMS ==========
	fmt.Println("\n4. DELETING ITEMS")
	
	fmt.Println("Before deletion - Fruits:", mp)
	delete(mp, "orange")
	fmt.Println("After deleting 'orange':", mp)
	
	fmt.Println("Before deletion - Numbers:", ap)
	delete(ap, 30)
	fmt.Println("After deleting key 30:", ap)
	
	// Safe to delete non-existent keys
	delete(mp, "dragonfruit") // No error
	fmt.Println("After deleting non-existent key:", mp)
	
	// ========== 5. ITERATING OVER MAPS ==========
	fmt.Println("\n5. ITERATING OVER MAPS")
	
	// Iterate with key and value
	fmt.Println("All capitals:")
	for country, capital := range capitals {
		fmt.Printf("  %s → %s\n", country, capital)
	}
	
	fmt.Println("All fruits and counts:")
	for fruit, count := range mp {
		fmt.Printf("  %s: %d\n", fruit, count)
	}
	
	// Iterate over keys only
	fmt.Println("Just country names:")
	for country := range capitals {
		fmt.Println("  ", country)
	}
	
	// ========== 6. MAP CHARACTERISTICS ==========
	fmt.Println("\n6. MAP CHARACTERISTICS")
	
	// Length of map
	fmt.Printf("Fruits map length: %d\n", len(mp))
	fmt.Printf("Capitals map length: %d\n", len(capitals))
	
	// Zero value behavior
	var nilMap map[string]int
	fmt.Println("Nil map is nil:", nilMap == nil)
	
	// This would cause runtime panic:
	// nilMap["key"] = 1 // ❌ Don't do this!
	
	// Maps are reference types
	reference := mp
	reference["grape"] = 12 // Affects original!
	fmt.Println("Original mp after reference modification:", mp)
	
	// ========== 7. ADVANCED MAP EXAMPLES ==========
	fmt.Println("\n7. ADVANCED MAP EXAMPLES")
	
	// Map of slices
	studentCourses := map[string][]string{
		"Alice":   {"Math", "Physics", "Chemistry"},
		"Bob":     {"History", "English"},
		"Charlie": {"Computer Science", "Math"},
	}
	
	// Add course to existing student
	studentCourses["Alice"] = append(studentCourses["Alice"], "Biology")
	
	fmt.Println("Student courses:")
	for student, courses := range studentCourses {
		fmt.Printf("  %s: %v\n", student, courses)
	}
	
	// Map of maps (nested)
	universityGrades := map[string]map[string]int{
		"Computer Science": {
			"Alice": 85,
			"Bob":   92,
		},
		"Mathematics": {
			"Charlie": 78,
			"Diana":   88,
		},
	}
	
	// Add new student to Computer Science
	universityGrades["Computer Science"]["Eve"] = 95
	
	fmt.Println("University grades by department:")
	for department, students := range universityGrades {
		fmt.Printf("  %s:\n", department)
		for student, grade := range students {
			fmt.Printf("    %s: %d\n", student, grade)
		}
	}
	
	// ========== 8. PRACTICAL USE CASES ==========
	fmt.Println("\n8. PRACTICAL USE CASES")
	
	// Word frequency counter
	text := "hello world hello go world go hello"
	words := strings.Fields(text)
	
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++ // Automatically handles zero value
	}
	fmt.Println("Word frequencies:", wordCount)
	
	// Phone book
	phoneBook := map[string]string{
		"Alice":   "555-1234",
		"Bob":     "555-5678",
		"Charlie": "555-9012",
	}
	
	// Lookup with existence check
	name := "Alice"
	if phone, exists := phoneBook[name]; exists {
		fmt.Printf("%s's phone number: %s\n", name, phone)
	}
	
	// Add new contact
	phoneBook["Diana"] = "555-3456"
	fmt.Println("All contacts:")
	for name, phone := range phoneBook {
		fmt.Printf("  %s: %s\n", name, phone)
	}
	
	// ========== 9. UTILITY FUNCTIONS ==========
	fmt.Println("\n9. UTILITY FUNCTIONS")
	
	// Get all keys
	keys := getKeys(mp)
	fmt.Println("All fruit keys:", keys)
	
	// Get all values
	values := getValues(mp)
	fmt.Println("All fruit counts:", values)
	
	// Merge maps
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}
	merged := mergeMaps(map1, map2)
	fmt.Println("Merged maps:", merged)
	
	// ========== 10. MAP WITH STRUCTS ==========
	fmt.Println("\n10. MAP WITH STRUCTS")
	
	type Student struct {
		Age    int
		Grade  string
		Courses []string
	}
	
	students := map[string]Student{
		"Alice": {
			Age:    23,
			Grade:  "A",
			Courses: []string{"Math", "Physics"},
		},
		"Bob": {
			Age:    22,
			Grade:  "B+",
			Courses: []string{"History", "English"},
		},
	}
	
	// Access struct fields
	fmt.Printf("Alice's grade: %s\n", students["Alice"].Grade)
	fmt.Printf("Bob's courses: %v\n", students["Bob"].Courses)
	
	// Add new student
	students["Charlie"] = Student{
		Age:    21,
		Grade:  "A-",
		Courses: []string{"Computer Science"},
	}
	
	fmt.Println("All students:")
	for name, student := range students {
		fmt.Printf("  %s: Age %d, Grade %s, Courses %v\n", 
			name, student.Age, student.Grade, student.Courses)
	}
	
	fmt.Println("\n=== END OF MAP GUIDE ===")
}

// Utility function to get all keys from a map
func getKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Utility function to get all values from a map
func getValues(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Utility function to merge two maps
func mergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)
	
	// Copy map1
	for k, v := range map1 {
		result[k] = v
	}
	
	// Copy map2 (overwrites duplicates from map1)
	for k, v := range map2 {
		result[k] = v
	}
	
	return result
}