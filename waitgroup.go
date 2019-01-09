package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // declare waitgroup

	wg.Add(1) // add goroutines to the waitgroup

	go func() {
		fmt.Println("Hello from Goroutine")

		wg.Done() // return control back to main function
	}()

	wg.Wait() // main function wait for controls for goroutine to complete
}
