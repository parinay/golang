package main

import "fmt"

const helloPrefix = "Hello, "

func helloWorld(name string) string {
	if name == "" {
		return helloPrefix + "World!"
	} else {
		return helloPrefix + name
	}
}
func main() {
	fmt.Println(helloWorld("World"))
}
