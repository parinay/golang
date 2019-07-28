package main

import (
	"fmt"
	"net"
)

func assertMethod() {
	addr, err := net.LookupHost("golangbotx.com")

	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("Operation timeout")
		} else if err.Temporary() {
			fmt.Println("Temporary error")
		} else {
			fmt.Println("Generic error: ", err)

		}
		return
	}
	fmt.Println("addr: ", addr)
}
