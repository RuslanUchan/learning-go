package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	iStr := "Mamam"

	i, err := strconv.Atoi(iStr)

	// Error handling
	// err is a value
	if err != nil {
		fmt.Println("Error happened:", err.Error())
	} else {
		fmt.Println(i)
	}

	r, err := Div(10, 0)

	if err != nil {
		fmt.Println("Error happened:", err.Error())
	} else {
		fmt.Println(r)
	}
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	result := x / y
	return result, nil
}
