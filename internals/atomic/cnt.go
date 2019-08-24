package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func cnt() {
	var count uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&count, 1)
				time.Sleep(time.Millisecond)
			}
		}()
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second)

	fmt.Println("Count:", atomic.LoadUint64(&count))
}
