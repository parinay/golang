package main

import (
	"fmt"
	"net/http"
)

func chanl() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkChanl(link, c)
	}

	// for msg := range c {
	// 	fmt.Println(msg)
	//}
	for index := 0; index < len(links); index++ {
		fmt.Println(<-c)
	}
}
func checkLinkChanl(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down"
	}
	c <- link + " is up!"
}
