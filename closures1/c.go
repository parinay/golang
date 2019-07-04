package main

import "fmt"

func appendstr() func(string) string {
	t := "hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
func main() {
	a := appendstr()
	b := appendstr()

	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))
	fmt.Println(a("Gooher"))
	fmt.Println(b("!"))
}
