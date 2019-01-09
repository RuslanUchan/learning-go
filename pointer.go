package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address

	fmt.Println(a)
	fmt.Println(*a) // value of memory address

	*a = 5
	// cant use a = 5 because type of a is *int which is an pointer to
	// memory address and 5 is an int,
	// so they didn't match and can't be operated on

	// to operate, we can make 'unpack' a values to be int
	// by doing *a
	fmt.Println(x) // value of x changed

	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)

	//
	// with strings
	//

	hello := "Hello"
	var strPointer *string

	strPointer = &hello

	// print memory address of hello
	fmt.Println(&hello)
	fmt.Println(strPointer)

	// value of the pointer
	fmt.Println("Value of the pointer", *strPointer)

	// w/o pointer
	fmt.Println(hello)
	change(hello)
	fmt.Println(hello)

	// w/ pointer
	fmt.Println(hello)
	changePointer(&hello)
	fmt.Println(hello)
}

// Pass by value
func change(v string) {
	v = v + " Golang"
}

// Pass by reference
func changePointer(v *string) {
	*v = *v + " Golang"
}
