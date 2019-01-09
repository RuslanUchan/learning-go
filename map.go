package main

import "fmt"

func main() {
	// Map = Key/Value pair
	// map[key]value
	var mapPerson map[int]string
	mapPerson = make(map[int]string)

	mapPerson[1] = "Hazard"
	mapPerson[2] = "Kovacic"
	mapPerson[3] = "Jorginho"
	mapPerson[4] = "Uchan"

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}

	// Check if key exists
	k, v := mapPerson[4]
	if !v {
		fmt.Println("Key not present in mapPerson")
	} else {
		fmt.Println(k)
	}
}
