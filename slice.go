package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50}

	mySlice = append(mySlice, 60)

	for i, v := range mySlice {
		fmt.Println(i, v)
	}
}
