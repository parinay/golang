package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// c := make(chan string)
	// go boring("Boring ", c)
	c := boring("boring!")
	k := boring("Interesting!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
		fmt.Printf("You say: %q\n", <-k)
	}
	fmt.Println("I am leaving cause on of you is boring")
}

func boring(msg string) <-chan string {
	// closeures
	c := make(chan string)
	// anonymous function
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}() //invocation of anonymous function
	return c
}
