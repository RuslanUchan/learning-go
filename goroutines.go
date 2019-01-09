package main

import (
	"fmt"
	"time"
)

func main() {
	go helloGo()

	time.Sleep(1 * time.Second)
	fmt.Println("This is main function!")
}

func helloGo() {
	fmt.Println("Hello Go!")
}
