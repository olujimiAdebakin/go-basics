// closure_examples.go
// -------------------------------------------------------------
// A CLOSURE in Go is a function that "remembers" variables
// from the scope where it was defined — even after that scope ends.
//
// Think of it as a function with its own mini-memory.
// Closures allow you to create functions that capture and use
// external variables, making them useful for callbacks,
// factories, and maintaining state.
// -------------------------------------------------------------

package main

import "fmt"

// -------------------------------------------------------------
// Example 1: Passing a closure as a function parameter
// -------------------------------------------------------------

// test2 takes a function that accepts an int and returns an int
func test2(myFunc func(int) int) {
	fmt.Println("test2 output:", myFunc(7))
}

func main() {

	// Anonymous function assigned to a variable (this is a closure)
	test := func(x int) int {
		return x * -1 // simple negation
	}

	fmt.Println("Example 1 -> test(8):", test(8)) // Calls closure directly
	test2(test)                                   // Passes closure as parameter

	// -------------------------------------------------------------
	// Example 2: Closure capturing outer variable (stateful)
	// -------------------------------------------------------------
	count := 0

	increment := func() int {
		count++              // closure captures variable "count"
		return count         // it remembers "count" between calls
	}

	fmt.Println("Example 2 ->", increment()) // 1
	fmt.Println("Example 2 ->", increment()) // 2
	fmt.Println("Example 2 ->", increment()) // 3

	// -------------------------------------------------------------
	// Example 3: Function returning a closure
	// -------------------------------------------------------------
	// multiplier returns a closure that multiplies by 'n'
	multiplier := func(n int) func(int) int {
		return func(x int) int {
			return x * n // closure captures 'n'
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("Example 3 -> double(5):", double(5)) // 10
	fmt.Println("Example 3 -> triple(5):", triple(5)) // 15

	// -------------------------------------------------------------
	// Example 4: Closure for filtering values in a slice
	// -------------------------------------------------------------
	filter := func(numbers []int, predicate func(int) bool) []int {
		var result []int
		for _, num := range numbers {
			if predicate(num) { // call the closure
				result = append(result, num)
			}
		}
		return result
	}

	numbers := []int{1, 2, 3, 4, 5, 6}

	// Create closures for even and odd checking
	isEven := func(x int) bool { return x%2 == 0 }
	isOdd := func(x int) bool { return x%2 != 0 }

	fmt.Println("Example 4 -> even numbers:", filter(numbers, isEven))
	fmt.Println("Example 4 -> odd numbers:", filter(numbers, isOdd))

	// -------------------------------------------------------------
	// Example 5: Closure used in a loop (demonstrating variable capture)
	// -------------------------------------------------------------
	funcs := []func() {
		var fns []func()
		for i := 1; i <= 3; i++ {
			val := i // create a local copy of i
			fns = append(fns, func() {
				fmt.Println("Example 5 -> closure captured value:", val)
			})
		}
		return fns
	}()

	for _, fn := range funcs {
		fn()
	}
}
