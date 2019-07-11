package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height int
}
type Circle struct {
	r float64
}

// pointer receivers
// method with same name area() for struct rect
func (r *rect) area() int {
	return r.width * r.height
}

// value receivers
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// value receivers
// method with same name area() for struct Circle

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perm() float64 {
	return 2 * math.Pi * c.r
}
func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	c := Circle{
		r: 12,
	}
	fmt.Printf("Area of circle is %f", c.area())
	fmt.Printf("Area of circle is %f", c.perm())
}
