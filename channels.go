package main

import (
	"fmt"
)

func main() {
	// Make channel
	done := make(chan bool)

	// pass channel to function invocations
	go helloGo(done)

	<-done
	fmt.Println("This is main function!")
}

// Pass channel to the function
func helloGo(done chan bool) {
	fmt.Println("Hello Go!")

	// assign channel to the function
	done <- true
}
