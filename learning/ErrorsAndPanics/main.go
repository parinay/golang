package main

import "fmt"

func main() {
	// Assert underlying struct type and get more information
	// from struct fields using func()
	assertError()

	// Assert underlying struct type and get more information
	// from struct fields using methods
	assertMethod()

	// match pattern
	assertCompare()

	// Custom Errors
	customErrors()

	// Custome Methods
	customMethods()

	//panic and recover
	a()
	fmt.Println("normally returned from a()")

}
