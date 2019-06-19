// This package is classic but useless
package main

import (
	"fmt"

	"github.com/parinay/golang/goTraining/first/firstlib"
)

func main() {
	// thenumber := rand.Intn(100) + 1
	thenumber := 100
	var guess int
	var i int
	// fmt.Println("Hello, World.")
	for i = 0; i < 5; {
		fmt.Print("Your guess? ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			println("Please enter a valid number")
			continue
		}
		i++
		message, done := firstlib.Checkguess(guess, thenumber)
		fmt.Println(message)
		if done {
			break
		}
	}
	if i >= 5 {
		fmt.Println("Out of retries ...")
	}
}
