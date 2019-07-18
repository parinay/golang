package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello World goroutine")
	// write to the boolean channel
	done <- true
}
func main() {
	fmt.Println("main function start")
	// channel of type boolean
	done := make(chan bool)
	go hello(done)

	// channel receiving data but do not store OR use the value
	<-done

	fmt.Println("main function - after channel receive")

	number := 589

	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)

	squares, cubes := <-sqrch, <-cubech

	fmt.Printf("The number is 589\nsquares:%d \ncubes:%d \nFinal Output(squares+cubes):%d, \n", squares, cubes, squares+cubes)

	// uni directional channel

	sendch := make(chan<- int)
	go sendData(sendch)

	// fmt.Printlsn(<-sendch)

	// chanl is bi-rectional in main()
	chanl := make(chan int)
	go sendData(chanl)
	fmt.Println(<-chanl)

	// buffered channels

	ch := make(chan string, 3)
	ch <- "abcd"
	ch <- "efgh"

	fmt.Println("Capacity is", cap(ch))
	fmt.Println("Length is", len(ch))

	fmt.Println("Read Value", <-ch)
	fmt.Println("Length is", len(ch))

	ch1 := make(chan int, 2)

	go write(ch1)
	time.Sleep(2 * time.Second)

	for v := range ch1 {
		fmt.Println("read value", v, "from ch1")
		time.Sleep(2 * time.Second)
	}
}
