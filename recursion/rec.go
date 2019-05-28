package main

import "fmt"

func fact1(n int) int {
	a := 1

	for i := 2; i <= n; i++ {
		a *= i
	}
	return a
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(7))
}
