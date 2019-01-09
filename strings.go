package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var myString = "Hello Golang :)"
	yo := "10"

	fmt.Println(myString)
	fmt.Println(len(myString))

	myStringUpper := strings.ToUpper(myString)
	fmt.Println(myStringUpper)

	resultContains := strings.Contains(myString, "Go")
	fmt.Println(resultContains)

	resultSplit := strings.Split(myString, " ") // Dynamic array
	for _, v := range resultSplit {
		fmt.Println(v)
	}

	// String conversion
	yow, err := strconv.Atoi(yo)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("type of yow is %T\n", yow)

	woy := strconv.Itoa(yow)
	fmt.Printf("type of woy is %T\n", woy)
}
