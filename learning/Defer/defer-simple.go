package main

import "fmt"

func deferSimple() {
	defer fmt.Print("world!")
	fmt.Print("Hello ")
}
