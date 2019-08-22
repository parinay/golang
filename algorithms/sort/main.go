package main

import "github.com/parinay/golang/algorithms/sort/global"

func main() {

	f := []global.Family{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}
	sliceStable(f)

	sorti(f)

	maps()
}
