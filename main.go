package main

import (
	"flag"
	"fmt"
	"sync"
	"test/adapter"
	"time"
)

type Request struct {
	ID int
	// Other request fields
}

type Response struct {
	ID   int
	Data string
	// Other response fields
}

// Function to process each request concurrently
func processRequest(req Request, wg *sync.WaitGroup, resultChan chan<- Response) {
	defer wg.Done()
	// Simulate processing time for the request
	time.Sleep(time.Second)
	// Simulate generating a custom report
	report := fmt.Sprintf("Report for request ID %d", req.ID)
	// Send the result through the channel
	resultChan <- Response{ID: req.ID, Data: report}
}

func main() {
	// Defining flags for tasking
	task := flag.String("task", "", "Specify the task to execute (adapter/concurrency)")
	flag.Parse()

	// Execute the correct task
	switch *task {
	case "adapter":
		adapterTask()
	case "concurrency":
		concurrencyTask()
	default:
		fmt.Println("Invalid task. Please specify 'adapter' or 'concurrency'")
	}
}

// Task #1: Implementation of the "Adapter" structural design pattern
func adapterTask() {
	var t string
	fmt.Println("What payment method will you use?")
	fmt.Scanln(&t)

	// sending the payment method
	paymentAdapter := adapter.Factory(t)
	// they will print Paid with Paypal or Stripe
	paymentAdapter.Pay()

}

// Task #2: Implementation of concurrency and parallelism
func concurrencyTask() {
	// Example incoming requests
	requests := []Request{{ID: 1}, {ID: 2}, {ID: 3}}

	var wg sync.WaitGroup
	resultChan := make(chan Response, len(requests))

	// Process each request concurrently
	for _, req := range requests {
		wg.Add(1)
		go processRequest(req, &wg, resultChan)
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Receive results from goroutines
	for res := range resultChan {
		fmt.Printf("Received report: %s\n", res.Data)
	}
}
