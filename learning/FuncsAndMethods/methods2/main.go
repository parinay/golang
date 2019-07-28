package main

import "fmt"

// type alias for built-int type int
type myInt int

// method with type alias as value receiver
func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	num1 := myInt(5)
	num2 := myInt(5)

	fmt.Println("Sum is ", num1.add(num2))
}
