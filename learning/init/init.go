// https://tutorialedge.net/golang/the-go-init-function/
package main

import "fmt"

var name string

func init() {
	fmt.Println("This will be called on main init")
	name = "Elliot"
}
func main() {
	fmt.Println("My wonderful Go program")
	fmt.Printf("Name is :%s \n", name)
}
