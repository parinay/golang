package main

import (
	"fmt"
)

// value receivers
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// value recivers - chnages are not reflected back
func (e Employee) changeName(newName string) {
	e.name = newName
}

// pointer reciver - changes are reflected back
func (e *Employee) changeSalary(newSalary int) {
	e.salary = newSalary
}

type Employee struct {
	name     string
	salary   int
	currency string
}

type address struct {
	city  string
	state string
}

type person struct {
	firstname string
	lastname  string
	address
}

// methods on anonymous fields of struct p, address
func (a address) fulladdress() {
	fmt.Printf("Full address: %s %s", a.city, a.state)
}

func main() {
	emp1 := Employee{
		name:     "Kaka",
		salary:   50000,
		currency: "USD ",
	}
	fmt.Println("Employee 1 - ", emp1)
	emp1.displaySalary()
	emp1.changeName("Mama")
	fmt.Println("Employee 1 (changeName)- ", emp1)
	emp1.changeSalary(0)
	fmt.Println("Employee 1 (changeSalary)- ", emp1)

	// methods of anonymous structs fields

	p := person{
		firstname: "Narayan",
		lastname:  "Murthy",
		address: address{
			city:  "B'lore",
			state: "K'taka",
		},
	}
	p.fulladdress()

}
