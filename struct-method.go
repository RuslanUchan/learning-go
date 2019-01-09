package main

import "fmt"

func main() {
	uchan := &Person{
		ID:   1,
		Name: "Uchan",
		Age:  21,
	}

	fmt.Println(uchan.GetName())
	uchan.ChangeName("Ruslan") // Change name
	fmt.Println(uchan.GetName())
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

// Method of the struct
// (p *Person) is a RECEIVER for the method
func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}
