package main

import "fmt"

func main() {
	Sum := 1
	for Sum <= 1000 {
		Sum += Sum
	}
	fmt.Println("Sum = Sum", Sum)
}
