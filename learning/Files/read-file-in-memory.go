package main

import (
	"fmt"
	"io/ioutil"
)

func readFileInMemory() {
	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		fmt.Println("File reading err", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
