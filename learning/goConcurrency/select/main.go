package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	// time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	// time.Sleep(3 * time.Second)
	ch <- "from server2"
}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	// radom case
	// comment the sleep in go routines so either
	// of the go routines can be executed
	time.Sleep(1 * time.Second)
	// select block until one of its cases is ready

	select {
	// blocks for 6 seconds
	case s1 := <-output1:
		fmt.Println(s1)
	// blocks for 3 seconds
	case s2 := <-output2:
		fmt.Println(s2)
	default:
		fmt.Println("Default case executed")
	}

}
