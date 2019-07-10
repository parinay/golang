package main

import (
	"fmt"
	"log"
)

func main() {
	//declare a channel
	var done chan bool
	done = make(chan bool)

	result := make(chan int)

	go doloop("path1", 10, done, result)
	go doloop("main", 15, done, result)

	var lastnumber int

	for i := 0; i < 2; i++ {
		<-done
		lastnumber = <-result
		fmt.Printf("Last number: %d\n", lastnumber)
	}
}
func doloop(name string, times int, finished chan bool, result chan int) {
	var lastnumber int
	defer log.Printf("Defer of %s \n", name)
	for i := 0; i < times; i++ {
		lastnumber = i
		log.Printf("%s: %d out of %d\n", name, i, times)
	}
	if finished != nil {
		finished <- true
	}
	result <- lastnumber

}
