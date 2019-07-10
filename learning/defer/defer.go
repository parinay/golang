package main

import "fmt"

func main() {
	defer fmt.Print("world!")
	fmt.Print("Hello ")
}
