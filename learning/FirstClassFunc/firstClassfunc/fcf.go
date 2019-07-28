package main

import "fmt"

func main() {
	fmt.Println("First Class functions")

	a := func(n string) {
		fmt.Printf("Hello to first class functions %s\n", n)
	}

	a("Gohers")
	fmt.Printf("Type of a :%T\n", a)
}
