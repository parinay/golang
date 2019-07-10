// https://medium.com/rungo/interfaces-in-go-ab1601159b3a

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	w float64
	h float64
}

func (r Rect) Area() float64 {
	return r.w * r.h
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.w + r.h)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func main() {
	var s Shape

	s = Rect{10, 3}

	fmt.Println("Value of s is: ", s)
	fmt.Printf("type of s is %T\n", s)
	fmt.Println("Area of rectangle s :", s.Area())
	fmt.Println("Perimeter of rectable s: ", s.Perimeter())

	s = Circle{10}
	fmt.Printf("type of s %T\n", s)
	fmt.Printf("Value of s is %v\n", s)
	fmt.Printf("Area of s is %0.2f\n", s.Area())
	fmt.Printf("Perimeter of s is %0.2f\n", s.Perimeter())
}
