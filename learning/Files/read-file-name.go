package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func readCmdLine() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	// fmt.Println("Value of fpath is - ", *fptr)
	data, err := ioutil.ReadFile(*fptr)

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
