package main

import "fmt"

/* CHANNEL BUFFER
* unbuffer channel ->
* done <- true // this is the sender
* <- done // this is the receiver
*
* This kind of channel DOES buffer | block process of other routines
* it's not really concurrent. It's blocking
 */

func main() {
	hello := make(chan string, 2)
	hello <- "Hello"
	hello <- "Golang"

	// need to close channel if we want to loop through it
	close(hello)

	// fmt.Println(<-hello)
	// fmt.Println(<-hello)

	// looping through channel
	for v := range hello {
		fmt.Println(v)
	}
}
