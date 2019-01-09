package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 5

	z := add(x, y)
	greetings := hello("Uchan")

	fmt.Println(z)
	fmt.Println(greetings)

	a := "Hello"
	b := "Golang"

	c, d := swap(a, b)
	fmt.Println(a, b)
	fmt.Println(c, d)
}

func add(x, y int) int {
	return x + y
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func swap(x, y string) (string, string) {
	return y, x
}
