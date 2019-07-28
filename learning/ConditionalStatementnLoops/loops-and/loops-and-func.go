package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := float64(1)
	for i := 1.0; i < x; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))

	}
	return z
}

func main() {
	n := 10

	for i := 0; i < n; i++ {
		sqrt_d := math.Sqrt(float64(i))
		sqrt_n := Sqrt(float64(i))

		fmt.Printf("math.Sqrt = %g", sqrt_d)
		fmt.Printf(" Sqrt = %g ", sqrt_n)
		fmt.Printf(" Difference: %g \n", math.Abs(sqrt_n-sqrt_d))
	}
}
