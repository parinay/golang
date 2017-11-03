package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d %v\n", len(s), cap(s), s)
}

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	var l []int
	fmt.Println(q)
	printSlice(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	q = q[0:4]
	fmt.Println(q)

	q = q[:2]
	fmt.Println(q)

	printSlice(q)

	q = q[0:0]
	printSlice(q)
	fmt.Println(q, len(q), cap(q))

	//	q = q[0:7]
	fmt.Println(l, len(l), cap(l))
	if l == nil {
		fmt.Println("nil!")
	}
}
