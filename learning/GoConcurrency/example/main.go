package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("### Sequential execution ###")
	fmt.Println(time.Now())
	seq()
	fmt.Println(time.Now())
	fmt.Println("### Concurrent execution using go routines ###")
	fmt.Println(time.Now())
	cc()
	fmt.Println(time.Now())
	fmt.Println("### Concurrent execution using channels ###")
	chanl()
	fmt.Println(time.Now())
}
