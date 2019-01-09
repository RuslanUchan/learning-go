package main

import "fmt"

func main() {
	// Traditional Conditionals
	t := true
	f := false

	if t == true && f == false {
		fmt.Println("If block executes")
	} else {
		fmt.Println("Else block executes")
	}

	// Switch Case
	x := 110

	switch x {
	case 50:
		fmt.Println("Case 50")
	case 100:
		fmt.Println("Case 100")
	default:
		fmt.Println("Default")
	}

	// Conditional Switch Case
	switch {
	case x < 100:
		fmt.Println("Less than 100")
	case x > 100:
		fmt.Println("More than 100")
	case x == 100:
		fmt.Printf("is 100!")
	default:
		fmt.Printf("Not even a number")
	}

	//
	// Type Switch
	//

	var y interface{}
	y = 5

	switch y.(type) {
	case int:
		fmt.Println("Integer type")
	case string:
		fmt.Println("String type")
	case float64:
		fmt.Println("Float type")
	default:
		fmt.Println("Type not found")
	}
}
