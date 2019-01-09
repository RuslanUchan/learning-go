package main

import (
	"errors"
	"fmt"
)

func main() {
	// Error handling stacked onto if else statements
	// inline error handling
	if x, err := Div(25, 0); err != nil {
		fmt.Println("Error happened:", err.Error())
	} else {
		fmt.Println(x)
	}
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	result := x / y
	return result, nil
}
