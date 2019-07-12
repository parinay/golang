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

// function with value argument

func area(r rect) {
	fmt.Printf("Area of rectangle - function results:%d \n", r.height*r.width)
}

// function with pointer argument
func perim(r *rect) {
	fmt.Println("Rectangle perimeter:", 2*r.width+2*r.height)
}

// method with value receivers
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

	fmt.Println("Rectangle area: ", r.area())
	fmt.Println("Rectangle perimeter:", r.perim())

	c := Circle{
		r: 12,
	}
	fmt.Printf("Area of circle is %f\n", c.area())
	fmt.Printf("Perimeter of circle is %f\n", c.perm())

	// function call with pointer
	// following will error out
	// area(rp)

	// function call with value argument
	area(r)

	// method call with pointer reciver
	rp := &r
	fmt.Println("Rectangle area (pointer): ", rp.area())
	fmt.Println("Rectangle perim (pointer): ", rp.perim())

	// function call with pointer param
	perim(rp)

	// calling pointer receiver with a value
	rp.perim()

}
