package main

import (
	"fmt"
)

func prints() {
	i := 0
	j := "golangbot.com"
	fmt.Print(i, " ", j) //no new line appended
	fmt.Println()
	fmt.Printf("%d %s", i, j) //no new line appended and formatting according to format verb
	fmt.Println()
	s := fmt.Sprintf("Welcome to %s", j) //string returned formatted according to format verb
	fmt.Println()
	fmt.Println(s)
}
