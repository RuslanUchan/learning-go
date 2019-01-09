package main

import "fmt"

func main() {
	// mapA := map[int]string
	mapPlayer := make(map[int]Player)

	mapPlayer[1] = Player{ID: 1, Name: "Hazard"}
	mapPlayer[2] = Player{ID: 2, Name: "Kovacic"}
	mapPlayer[3] = Player{ID: 3, Name: "Jorginho"}

	for _, v := range mapPlayer {
		fmt.Println(v.Name)
	}
}

type Player struct {
	ID   int
	Name string
}
