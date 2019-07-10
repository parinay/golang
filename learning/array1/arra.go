package main

import "fmt"

func main() {
	package main

import (
	"fmt"
)

func main() {
	sl1 := make([]int, 4, 4)
	fmt.Println("sl1: ",sl1)

	sl1 = append(sl1, 42, 44)
	fmt.Println("sl1: ",sl1)
	length := len(sl1)
	cap := cap(sl1)

	fmt.Println("Length sl1: ",length)
	fmt.Println("capacity ",cap)
	sl2 := []string{"a", "b", "c"}
	fmt.Println("sl2: ",sl2)

	length2 := len(sl2)
	cap2 := cap(sl2)
	fmt.Println("Length sl1: ",length2)
	fmt.Println("capacity ",cap2)
}
}
