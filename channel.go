// Simple Explanation:
// Channels are like pipes or conveyor belts that allow different goroutines to communicate and synchronize with each other safely.

// Basic Concept:
// Channel = A typed conduit for sending/receiving data between goroutines

// Like a tube where one goroutine puts data in and another takes it out

// Thread-safe - built-in synchronization


package main

import (
	"fmt"
	"time"
)

// ========== BASIC CHANNEL OPERATIONS ==========

func basicChannelExamples() {
	fmt.Println("=== BASIC CHANNEL EXAMPLES ===")
	
	// Example 1: Creating channels
	fmt.Println("🎯 Example 1: Channel Creation")
	
	// Unbuffered channel (size 0) - synchronous communication
	// Sender and receiver must be ready at the same time
	ch1 := make(chan int)
	fmt.Printf("Unbuffered channel: %T\n", ch1)
	
	// Buffered channel (size 3) - asynchronous communication  
	// Can store multiple values before blocking
	ch2 := make(chan string, 3)
	fmt.Printf("Buffered channel: %T\n", ch2)
	
	// Example 2: Basic send/receive
	fmt.Println("\n🎯 Example 2: Basic Send/Receive")
	
	// Start a goroutine to send data
	go func() {
		ch1 <- 42 // Send value to channel (blocks until someone receives)
		fmt.Println("Sent 42 to channel")
	}()
	
	// Receive value from channel (blocks until someone sends)
	value := <-ch1 
	fmt.Printf("Received: %d\n", value)
	
	// Example 3: Buffered channel operations
	fmt.Println("\n🎯 Example 3: Buffered Channels")
	
	// Send multiple values without blocking (up to buffer size)
	ch2 <- "Hello"
	ch2 <- "World"
	ch2 <- "!"
	fmt.Println("Sent 3 messages to buffered channel")
	
	// Receive them in FIFO order (First In, First Out)
	fmt.Println("Received:", <-ch2)
	fmt.Println("Received:", <-ch2)
	fmt.Println("Received:", <-ch2)
}

// ========== CHANNEL TYPES ==========

func channelTypes() {
	fmt.Println("\n=== CHANNEL TYPES ===")
	
	// Unbuffered channel (synchronous) - direct handoff
	fmt.Println("🔁 UNBUFFERED CHANNEL (Synchronous)")
	unbuffered := make(chan int)
	
	go func() {
		fmt.Println("Goroutine: Waiting to send...")
		unbuffered <- 100 // Blocks until someone receives
		fmt.Println("Goroutine: Send completed!")
	}()
	
	time.Sleep(1 * time.Second) // Simulate some work
	fmt.Println("Main: Receiving...")
	value := <-unbuffered // Now the send can complete
	fmt.Printf("Main: Received %d\n", value)
	
	// Buffered channel (asynchronous) - like a mailbox
	fmt.Println("\n📦 BUFFERED CHANNEL (Asynchronous)")
	buffered := make(chan int, 2) // Can hold 2 values
	
	buffered <- 1 // Doesn't block (buffer has space)
	buffered <- 2 // Doesn't block
	fmt.Println("Sent 2 values without blocking")
	
	// Receive values
	fmt.Println("Received:", <-buffered)
	fmt.Println("Received:", <-buffered)
}

// ========== CHANNEL DIRECTIONS ==========

// Channel as parameter with direction for type safety
// chan<- int means this function can only SEND to the channel
func sender(ch chan<- int) { 
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch) // Important: close when done sending to signal no more data
}

// <-chan int means this function can only RECEIVE from the channel
func receiver(ch <-chan int) { 
	// range over channel automatically stops when channel is closed
	for value := range ch { 
		fmt.Printf("Received: %d\n", value)
		time.Sleep(500 * time.Millisecond) // Simulate processing time
	}
}

func channelDirections() {
	fmt.Println("\n=== CHANNEL DIRECTIONS ===")
	
	ch := make(chan int, 2) // Create a buffered channel
	
	go sender(ch) // Start sender goroutine
	receiver(ch)  // Run receiver in main goroutine
}

// ========== CHANNEL SYNCHRONIZATION ==========

// Worker function that processes jobs from input channel and sends results to output channel
func worker(id int, jobs <-chan int, results chan<- int) {
	// Process jobs until jobs channel is closed
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(1 * time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send result back
	}
}

func workerPoolExample() {
	fmt.Println("\n=== WORKER POOL PATTERN ===")
	
	const numJobs = 5
	const numWorkers = 3
	
	// Create channels for jobs and results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}
	
	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Important: close to signal no more jobs
	
	// Collect results from all jobs
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Result for job %d: %d\n", r, result)
	}
}

// ========== SELECT STATEMENT ==========

// Select allows waiting on multiple channel operations simultaneously
func selectExample() {
	fmt.Println("\n=== SELECT STATEMENT ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Goroutine that sends to ch1 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()
	
	// Goroutine that sends to ch2 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()
	
	// Select will choose the first channel that becomes ready
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1) // This will fire first (after 1 second)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2) // This will fire second (after 2 seconds)
		case <-time.After(3 * time.Second): // Timeout case
			fmt.Println("Timeout!")
			return
		}
	}
}

// ========== CHANNEL CLOSING ==========

func closingChannels() {
	fmt.Println("\n=== CHANNEL CLOSING ===")
	
	ch := make(chan int, 3)
	
	// Send some values
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // Close channel - no more sends allowed after this
	
	// Receiving from closed channel
	fmt.Println("Receiving from closed channel:")
	for i := 0; i < 4; i++ {
		value, ok := <-ch
		if ok {
			fmt.Printf("Received: %d\n", value) // Can still receive buffered values
		} else {
			fmt.Println("Channel closed!") // No more values to receive
		}
	}
	
	// Using range with closed channel
	fmt.Println("\nUsing range with channel:")
	ch2 := make(chan int, 2)
	ch2 <- 10
	ch2 <- 20
	close(ch2) // Must close channel before ranging over it
	
	// range automatically stops when channel is closed and drained
	for value := range ch2 {
		fmt.Printf("Range received: %d\n", value)
	}
}

// ========== REAL-WORLD EXAMPLES ==========

// Producer-Consumer Pattern: One goroutine produces data, another consumes it
func producerConsumer() {
	fmt.Println("\n=== PRODUCER-CONSUMER PATTERN ===")
	
	jobs := make(chan int, 5)
	done := make(chan bool) // Signal channel to indicate completion
	
	// Producer goroutine - generates data
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
			fmt.Printf("Produced: %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(jobs) // Signal that no more data will be produced
	}()
	
	// Consumer goroutine - processes data
	go func() {
		for job := range jobs { // Automatically stops when jobs is closed
			fmt.Printf("Consumed: %d\n", job)
			time.Sleep(1 * time.Second) // Simulate processing time
		}
		done <- true // Signal that all jobs are processed
	}()
	
	<-done // Wait for consumer to finish processing all jobs
	fmt.Println("All jobs processed!")
}

// Fan-out, Fan-in Pattern: Multiple workers process data, results are combined
func fanOutFanIn() {
	fmt.Println("\n=== FAN-OUT, FAN-IN PATTERN ===")
	
	input := make(chan int)   // Channel for input data
	output := make(chan int)  // Channel for output results
	
	// Start multiple workers (fan-out) - parallel processing
	for i := 1; i <= 3; i++ {
		go workerFan(i, input, output)
	}
	
	// Send inputs to workers
	go func() {
		for i := 1; i <= 6; i++ {
			input <- i
		}
		close(input) // Signal no more input
	}()
	
	// Collect outputs from all workers (fan-in)
	go func() {
		for i := 1; i <= 6; i++ {
			result := <-output
			fmt.Printf("Final result: %d\n", result)
		}
		close(output)
	}()
	
	time.Sleep(3 * time.Second) // Wait for processing to complete
}

// Worker function for fan-out pattern
func workerFan(id int, input <-chan int, output chan<- int) {
	for num := range input {
		fmt.Printf("Worker %d processing: %d\n", id, num)
		time.Sleep(500 * time.Millisecond) // Simulate work
		output <- num * num // Send result (square of input)
	}
}

// ========== CHANNEL WITH CONTEXT ==========

// Using channels for graceful shutdown/control
func contextWithChannels() {
	fmt.Println("\n=== CHANNEL WITH CONTEXT ===")
	
	messages := make(chan string) // Channel for work items
	quit := make(chan bool)       // Channel for shutdown signal
	
	// Worker that can be stopped gracefully
	go func() {
		for {
			select {
			case msg := <-messages:
				fmt.Printf("Processing: %s\n", msg) // Process incoming messages
			case <-quit:
				fmt.Println("Worker stopping...") // Received shutdown signal
				return // Exit goroutine
			}
		}
	}()
	
	// Send some messages to worker
	messages <- "Hello"
	messages <- "World"
	
	// Stop the worker by sending quit signal
	quit <- true
	time.Sleep(100 * time.Millisecond) // Give worker time to clean up
	fmt.Println("Worker stopped!")
}

// ========== ERROR HANDLING WITH CHANNELS ==========

// Using channels to communicate both results and errors
func errorHandling() {
	fmt.Println("\n=== ERROR HANDLING ===")
	
	// Result struct to carry both value and error
	type Result struct {
		Value int
		Error error
	}
	
	results := make(chan Result)
	
	// Goroutine that succeeds
	go func() {
		time.Sleep(1 * time.Second)
		results <- Result{Value: 42, Error: nil} // Success case
	}()
	
	// Goroutine that fails
	go func() {
		time.Sleep(2 * time.Second)
		results <- Result{Value: 0, Error: fmt.Errorf("something went wrong")} // Error case
	}()
	
	// Process results from both goroutines
	for i := 0; i < 2; i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error) // Handle error
		} else {
			fmt.Printf("Success: %d\n", result.Value) // Handle success
		}
	}
}

// ========== MAIN FUNCTION ==========

func main() {
	fmt.Println("🎯 CHANNELS IN GO - COMPLETE GUIDE")
	fmt.Println("===================================")
	
	basicChannelExamples()    // Basic send/receive operations
	channelTypes()           // Unbuffered vs buffered channels
	channelDirections()      // Send-only and receive-only channels
	workerPoolExample()      // Worker pool pattern
	selectExample()          // Select statement for multiple channels
	closingChannels()        // Proper channel closing
	producerConsumer()       // Producer-consumer pattern
	fanOutFanIn()           // Fan-out, fan-in pattern
	contextWithChannels()    // Graceful shutdown with channels
	errorHandling()         // Error handling patterns
	
	fmt.Println("\n=== CHANNEL GUIDE COMPLETE ===")
}