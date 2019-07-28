package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started Go Routine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Go Routine %d ended\n", i)
	//Decrements the WaitGroup Counter by 1
	wg.Done()
}
func main() {
	no := 3

	// WaitGroup{}
	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		//Add 1 to WaitGorup counter
		wg.Add(1)
		// its imp to pass address of wg here, else each goroutine
		// will have its own wait group and main() wont be notified
		// when they finish executing
		go process(i, &wg)
	}

	//WaitGroup wait() method blocks the GoRoutine in which its
	// called, here main() untill the counter becomes zero
	wg.Wait()

	fmt.Println("All go routines finished executing.")
}
