package main

import (
	"fmt"
	"net/http"
)

func seq() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	fmt.Println("http get on %", link)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		return
	}
}
