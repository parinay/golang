package main

import (
	"fmt"
	"sync"
)

// global var, should be avoided as much as possible
var x = 0
var y = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {

	// lock() before updating the val of x
	m.Lock()
	x += 1
	// unlock() after updating the val of x,
	// making sure at any time only one go rotine
	// can access x
	m.Unlock()
	// Done() waitgroup
	wg.Done()
}
func incrementC(wg *sync.WaitGroup, ch chan bool) {
	// this buffered channel is capacity of 1
	// passing true, we enforce other routines to wait
	ch <- true
	// increment
	y += 1
	// waiting to read the value from channel after increment
	<-ch
	// Done() waitgroup
	wg.Done()

}
func main() {

	// Using mutex
	// zero valued WaitGroup type var
	var w sync.WaitGroup

	// zero valued sync.Mutex type var
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		// Add to the waitgroup
		w.Add(1)
		// spawn a go routine passing address of waitgroup var &&
		// address of sync.Mutex type to take care of race and
		// critical section
		go increment(&w, &m)
	}
	// wait for all wait group to finish
	w.Wait()

	// Everytime we run this we might receive new value for
	// x if Mutex is not used
	fmt.Println("final value of x", x)

	// Using buffered channel of capacity 1
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		// Add to the waitgroup
		w.Add(1)
		// spawn a go routine passing address of waitgroup var &&
		// buffered channel to take care of race and
		// critical section
		go incrementC(&w, ch)
	}
	// wait for all wait group to finish
	w.Wait()

	fmt.Println("final value of y", y)

}
