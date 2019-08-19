package main

import (
	"fmt"
	"net/http"
)

func cc() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	for _, link := range links {
		go checkLinkCC(link)
	}
}

func checkLinkCC(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		return
	}
}
