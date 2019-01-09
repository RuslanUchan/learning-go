package main

import "fmt"

func main() {
	defer fmt.Println("This is the last execution 2")

	defer fmt.Println("This is the last execution 1")

	fmt.Println("This is the first execution")
}
