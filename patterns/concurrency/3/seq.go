package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Message struct for squencing
type Message struct {
	str  string
	wait chan bool
}

func boring(msg string) <-chan Message {
	// closeures
	c := make(chan Message)
	waitForIt := make(chan bool)
	// anonymous function
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}() //invocation of anonymous function
	return c
}
func fanIn(input1, input2 <-chan Message) <-chan Message {

	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg1.wait <- true
	}
	fmt.Println("You're both boring, I am leaving.")
}
