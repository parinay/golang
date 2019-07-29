package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// Create a custom error message with errors.New()
		// return 0, errors.New("Area calculation failed, radius is less than zero")
		// Using fmt.Errorf()
		// return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)

		// Define a struct that uses Error{} interface to create custom error messages
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}
func customErrors() {
	radius := -20.0

	area, err := circleArea(radius)

	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("Area of circle is %0.2f", area)
}
