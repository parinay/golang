package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello World goroutine")
	// write to the boolean channel
	done <- true
}
func main() {
	// channel of type boolean
	done := make(chan bool)
	go hello(done)

	// channel receiving data but done not store OR use the value
	<-done

	fmt.Println("main function")
}
