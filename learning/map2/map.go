package main

import (
	"fmt"
)

func main() {

	var personSalary map[string]int

	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}
	personSalary["steve"] = 12000
	personSalary["jamie"] = 10000
	personSalary["mike"] = 15000

	fmt.Println("person salary map contents:", personSalary)

	employee := "jamie"

	fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"])

	newEmp := "joe"
	value, ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] =%d\n", key, value)
	}
	// blank identifier for key
	for _, value := range personSalary {
		fmt.Printf("personSalary = %d\n", value)
	}
	// blank identifier for value
	for key, _ := range personSalary {
		fmt.Printf("personSalary[%s]\n", key)
	}
	//Delete items

	fmt.Println("map before deletetion is", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletetion is", personSalary)

	// length of the map
	personSalary["mike"] = 16000
	personSalary["steve"] = 12000

	fmt.Println("current map is", personSalary)
	fmt.Println("length of map personSalary is", len(personSalary))

	// Maps are references

	newPersonSalary := personSalary

	newPersonSalary["selvam"] = 90000
	fmt.Println("Person salary changed", personSalary)

}
