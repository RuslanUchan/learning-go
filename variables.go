package main

import "fmt"

const A string = "This is a constant"

func main() {
	// Static type declaration
	var x int = 10
	var y float64 = 5.5

	fmt.Printf("type of %d is %T\n", x, x)
	fmt.Printf("type of %f is %T\n", y, y)

	// Dynamic type declaration
	z := "Uchan"
	fmt.Println(z)
	fmt.Printf("type of z is %T\n", z)

	// Constant //
	fmt.Println(A)

	const b int = 10

	fmt.Println(b)

	// This would trigger an error
	// b = 20
}
