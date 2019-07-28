package main

import "fmt"

func digits(number int, dg chan int) {
	defer close(dg)
	for number != 0 {
		digit := number % 10
		dg <- digit
		number /= 10
	}
}

func calcSquares(number int, sq chan int) {
	sum := 0

	dg := make(chan int)

	go digits(number, dg)
	// for range form of for loop to receive values from channel
	for digit := range dg {
		sum += digit * digit
	}
	sq <- sum
}

func calcCubes(number int, qb chan int) {
	sum := 0

	dg := make(chan int)
	go digits(number, dg)
	// for range form of for loop to receive values from channel
	for digit := range dg {
		sum += digit * digit * digit
	}
	qb <- sum
}

// channel with send only channel
func sendData(sendch chan<- int) {
	sendch <- 10
}

func write(ch chan int) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, "to ch")
	}
}
