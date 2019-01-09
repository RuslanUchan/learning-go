package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	x := 10
	var y int

	for y < x {
		y++
		if y == 5 {
			fmt.Println("Skipped 5")
			continue
		}
		fmt.Println(y)

		if y == 7 {
			fmt.Println("Break at 7")
			break
		}
	}
}
