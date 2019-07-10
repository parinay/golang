package main

import (
	"fmt"
)

func listNumbers(list []*int) {
	var number int

	// Loop from 0 to 9
	for i := 0; i < 10; i++ {

		number = i
		list = append(list, &number)
		fmt.Printf("%d \n", list[i])
		fmt.Printf("%d \n", *(list[i]))
	}

}
func main() {
	// Declare a list/slice of int pointers
	listOfNumberInt := []*int{}

	listNumbers(listOfNumberInt)
	/*
		var number int

		// Loop from 0 to 9
		for i := 0; i < 10; i++ {

			number = i
			listOfNumberInt = append(listOfNumberInt, &number)
			fmt.Printf("%d \n", listOfNumberInt[i])
			fmt.Printf("%d \n", *(listOfNumberInt[i]))
		}
	*/

}
