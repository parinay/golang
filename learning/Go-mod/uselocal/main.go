package main

import (
	"fmt"

	local "github.com/parinay/golang/learning/gomod/local"
)

func main() {
	fmt.Println("Using mod")
	fmt.Printf("Double of 23 is %d\n", local.Db1X(23))
}
