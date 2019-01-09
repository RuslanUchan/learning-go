package main

import "fmt"

func main() {
	f := func(v string) bool {
		return v == "Golang"
	}

	// Higher order function in action !
	resultTrue := match("Golang", f)
	resultFalse := match("Go", f)

	fmt.Println(result)
}

// f function is an argument
// match function is higher order function
func match(v string, f func(string) bool) bool {
	if f(v) {
		return true
	}
	return false
}
