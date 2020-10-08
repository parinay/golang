package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Var - Web/Image/Video a fake framework for search
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

// Result list to append search results
type Result string

// Search function prototype
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// GoogleSeq - search a query string across web, image, video and collate, sequentially
func GoogleSeq(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// GoogleParallel - search a query string across web, image, video and collate, parallelly
func GoogleParallel(query string) (results []Result) {
	c := make(chan Result)

	// fan-in pattern - get the data back on same channel
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	// timeout pattern of concurrency
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

// First replicates the servers, sending requests to multple replicas and use first response
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }

	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}
func main() {
	rand.Seed(time.Now().UnixNano())

	// 1.
	fmt.Println("Seqeuntial Search")
	start := time.Now()
	results := GoogleSeq("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
	// 2.
	fmt.Println("Parallel Search")
	start = time.Now()
	results = GoogleParallel("golang")
	elapsed = time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)

	// 3.
	fmt.Println("Using Multiple servers for Search")
	start = time.Now()
	result := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
	elapsed = time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

}
