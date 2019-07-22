package main

import "fmt"

type bill struct {
	citizen bool
	salary  int
	tax     float64
}
type ana struct {
	citizen bool
	salary  int
	tax     float64
}

func one() {
	var b bill
	var a ana

	e1 := struct {
		citizen bool
		salary  int
		tax     float64
	}{
		citizen: true,
		salary:  1000,
		tax:     101.4,
	}
	// implicit conversion - will fail as go doesn't support it
	//b = a

	// unless user knowingly/explicitly goes for conversion
	b = bill(a)
	fmt.Println(b, a)

	// e1 is a literal thus following works.
	b = e1
	fmt.Println(a, b, e1)

	// constant and conversions
	d := 1 + 3.4 + 8.0

	const x int = 5
	const y float64 = 5.4

	// c := x + y fails
	c := float64(x) + y

	fmt.Println(c, d)
}
