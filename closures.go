package main

import "fmt"

func main() {
	nextValue := genNumber()

	fmt.Println(nextValue()) // Invoked
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	hi := hello("Uchan")
	fmt.Println(hi("Golang"))
	fmt.Println(hi("Python"))
}

func genNumber() func() int {
	x := 0

	// Inner function return value must match
	// the outer function
	return func() int {
		x++
		return x
	}
}

func hello(name string) func(string) string {
	return func(things string) string {
		return fmt.Sprintf("%s code %s", name, things)
	}
}
