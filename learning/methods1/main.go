package main

import (
	"fmt"
)

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Employee struct {
	name     string
	salary   int
	currency string
}

func main() {
	emp1 := Employee{
		name:     "Kaka",
		salary:   50000,
		currency: "USD ",
	}
	emp1.displaySalary()
}
