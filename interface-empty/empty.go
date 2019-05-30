package main

import "fmt"

type MyString string

type Rect struct {
	w float64
	h float64
}

func explain(i interface{}) {
	fmt.Printf("Value given to explain function is of type '%T' and value '%v'\n", i, i)
}

func main() {
	ms := MyString("HEllO WORLD!")
	r := Rect{5.5, 4.5}

	explain(ms)
	explain(r)
}
