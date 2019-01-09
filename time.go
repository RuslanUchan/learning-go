package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())

	fmt.Println("=====\n")

	// Parsing
	timeString := "March 19, 2002"
	form := "September 11, 2001"

	resultTime, err := time.Parse(form, timeString)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultTime)

	// Equal and Dates
	t1 := time.Date(2011, 9, 3, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2011, 9, 3, 12, 0, 0, 0, time.UTC)

	eq := t1.Equal(t2)

	fmt.Println(eq)

	// Formatting
	layout := "2009-07-01"
	timeNow := time.Now()
	resultString := timeNow.Format(layout)

	fmt.Println(resultString)
}
