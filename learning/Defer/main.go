package main

import "fmt"

func main() {
	// simple defer fmt.Println()
	deferSimple()
	fmt.Println()
	// defer a method
	deferMethod()

	// defer a func ()
	deferFunc()

	//practical defer use
	deferPract()
}
