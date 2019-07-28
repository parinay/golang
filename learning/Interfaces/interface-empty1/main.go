package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am string and valye is %s\n", i.(string))
	case int:
		fmt.Printf("I am string and valye is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type")

	}
}
func main() {
	findType("naveen")
	findType(88)
	findType(88.99)
}
