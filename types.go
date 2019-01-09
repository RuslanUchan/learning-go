package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

func concat(a, b string) (string, string) {
	return a, b
}

func main() {
	// Unpacking - shorthand
	// num1, num2 := 9.5, 5.6 // default to float64

	word1, word2 := "Hello", "World!"

	fmt.Println(concat(word1, word2))

	// Type conversion
	var a = 10
	var b = float64(a)
	x := float64(a)

	fmt.Println(add(x, b))
}
