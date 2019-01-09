package main

import "fmt"

type car struct {
	// Car game example

	/*
	   lower_case attribute is PRIVATE
	   upper_case attribute is PUBLIC
	*/
	gasPedal      uint16 // min 0 to 65k
	brakePedal    uint16
	steeringWheel int16
	topSpeedKMH   float64
}

func main() {
	aCar := car{gasPedal: 10000,
		brakePedal:    0,
		steeringWheel: 10000,
		topSpeedKMH:   250.0}

	fmt.Println(a_car.gasPedal)
}
