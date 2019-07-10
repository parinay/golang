package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "India"
	fmt.Println("vim-go")
	fmt.Println("Hello", World)
	fmt.Println("Hello", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
