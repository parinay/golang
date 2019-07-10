package main

import (
	"fmt"
)

// Normal go struct declaration

type Employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
}
type EmpData struct {
	int
	string
}

type Pin struct {
	int
}
type Person struct {
	name    string
	age     int
	address Address
	Pin
}
type Address struct {
	city, state string
}

func main() {

	// Normal go struct init
	emp1 := Employee{
		firstname: "sam",
		lastname:  "k",
		age:       40,
		salary:    10000,
	}

	emp2 := Employee{"Thomas", "L", 45, 50000}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// anonymous struct
	emp3 := struct {
		firstname, lastname string
		age, salary         int
	}{
		firstname: "Naveen",
		lastname:  "R",
		age:       21,
		salary:    1500000,
	}

	fmt.Println("Employee 3", emp3)

	// zero value of struct

	var emp4 Employee
	fmt.Println("Employee 4", emp4)

	// Initalize some of the values

	emp5 := Employee{
		firstname: "John",
		lastname:  "paul",
	}

	fmt.Println("Employee 5", emp5)

	// Accessing individual fields

	emp6 := Employee{"Thomas", "anderson", 55, 60000}

	fmt.Println("Employee 6 -First Name:", emp6.firstname)
	fmt.Println("Employee 6 -Last Name:", emp6.lastname)

	// Assign later

	var emp7 Employee

	emp7.firstname = "jackie"
	emp7.lastname = "chan"

	fmt.Println("Employee 7 -Firstname:", emp7.firstname)

	// pointers to struct

	emp8 := &Employee{"Anand", "N", 10, 100000}

	fmt.Println("Employee 8 -First name:", (*emp8).firstname)
	fmt.Println("Employee 8 -Salary:", emp8.salary)

	// structs with anon fields

	ed := EmpData{1000, "Machester"}

	fmt.Println("Employee Data 1", ed)

	var ed1 EmpData

	ed1.int = 1001
	ed1.string = "Rajaji Nagar"

	fmt.Println("Employee Data 2", ed1)

	// Nested structs
	// promoted field

	var p Person

	p.name = "Raghu"
	p.age = 35
	p.address.city = "Pune"
	p.address.state = "maha"

	p.int = 50058
	fmt.Println("Person 1", p)

}
