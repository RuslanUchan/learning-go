package main

import "fmt"

// Constant
const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

// Struct
type car struct {
	gasPedal      uint16 // min 0 to 65k
	brakePedal    uint16
	steeringWheel int16
	topSpeedKMH   float64
}

// Struct methods
// Value Receiver
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax / kmh_multiple)
}

// Pointer Receiver
func (c *car) newTopSpeed(newspeed float64) {
	c.topSpeedKMH = newspeed
}

// Normal function
func newerTopSpeed(c car, speed float64) car {
	c.topSpeedKMH = speed
	return c
}

// Main function
func main() {
	aCar := car{gasPedal: 10000,
		brakePedal:    0,
		steeringWheel: 10000,
		topSpeedKMH:   250.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())

	aCar.new_top_speed(500)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())

	// redefine car
	aCar = newer_top_speed(aCar, 1500)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}
