package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	random := rand.Intn(999)
	data <- random
	wg.Done()
}

func consume(data chan int, done chan bool) {
	fd, err := os.Create("concurrent")

	if err != nil {
		fmt.Println("Create file failed:", err)
		return
	}

	for datai := range data {
		_, err = fmt.Fprintln(fd, datai)
		if err != nil {
			fmt.Println("File write failed:", err)
			fd.Close()
			done <- false
			return
		}
	}
	err = fd.Close()
	if err != nil {
		fmt.Println("File close failed:", err)
		done <- false
		return
	}

	done <- true
}

func concurrent() {
	data := make(chan int)
	done := make(chan bool)

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}

	go consume(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()

	d := <-done

	if d == true {
		fmt.Println("Concurrent file written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
