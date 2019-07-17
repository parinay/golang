package main

import "fmt"

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
}
