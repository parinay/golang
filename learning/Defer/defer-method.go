package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) fullname() {
	fmt.Printf("%s %s ", p.firstName, p.lastName)
}

func deferMethod() {
	p := person{
		firstName: "John",
		lastName:  "s",
	}
	defer p.fullname()
	fmt.Printf("Welcome ")
}
