package main

import (
	"fmt"

	computer "github.com/parinay/golang/learning/structs-2/computer/structs"
	computer1 "github.com/parinay/golang/learning/structs-2/computer1/strings"
)

func main() {
	var spec computer.Spec

	spec.Maker = "apple"
	spec.Price = 50000

	var slen computer1.StrLen
	slen.Length = 100

	fmt.Println("Spec:", spec)
	fmt.Println("Slen:", slen)
}
