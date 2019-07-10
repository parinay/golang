package main

import (
	"fmt"

	computer "github.com/parinay/golang/learning/structs-2/computer/structs"
)

func main() {
	var spec computer.Spec

	spec.Maker = "apple"
	spec.Price = 50000

	fmt.Println("Spec:", spec)
}
