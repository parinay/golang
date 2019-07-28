package main

import "fmt"

type Shape interface {
	Area() float64
}

type Obj interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}

	var s Shape = c

	var o Obj = c

	fmt.Println("Volume of s of interface type Shape is ", s.Area())
	fmt.Println("Area of o of interface type Obj is", o.Volume())

	var st Shape = Cube{3}
	c1 := st.(Cube)
	fmt.Println("Area of c of type Cube is", c1.Area())
	fmt.Println("Volume of c of type Cube is", c1.Volume())

}
