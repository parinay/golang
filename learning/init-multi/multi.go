package main

import "fmt"

var initCount int

func init() {
	fmt.Println("Called First in order of Declaration")
	initCount++
}
func init() {
	fmt.Println("Called second in order of Declaration")
	initCount++
}
func main() {
	fmt.Println("Does nothing of any Significance")
	fmt.Printf("Init counter: %d\n", initCount)
}
