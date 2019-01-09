package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang Select Channel")

	// Illustrate a delay
	httpRequest1 := process(1, 3)
	httpRequest2 := process(2, 5)
	httpRequest3 := process(3, 1)

	for i := 1; i <= 3; i++ {
		select {
		case c1 := <-httpRequest1:
			fmt.Println(c1)
		case c2 := <-httpRequest2:
			fmt.Println(c2)
		case c3 := <-httpRequest3:
			fmt.Println(c3)
		}
	}
}

func process(id int, delay time.Duration) <-chan string {
	/*
		output: channel with specific worker id
		delay: operation of a certain process (database, connection, etc)
	*/

	output := make(chan string) // make channels
	go func() {
		// create goroutines that affected by time delay of a certain process
		// and accept input from the certain process
		time.Sleep(time.Second * delay)

		output <- fmt.Sprintf("Process %d", id)
	}()

	return output
}
