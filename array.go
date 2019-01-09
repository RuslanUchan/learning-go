package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arrStr := [5]string{"Hi", "Hello", "Howdy", "Hey", "Yo"}
	fmt.Println(arrStr[4])
}
