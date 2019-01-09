package main

import "fmt"

func main() {
	uchan := Person{
		ID:   1,
		Name: "Uchan",
		Age:  21,
	}

	printPerson(uchan)
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func printPerson(p Person) {
	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
