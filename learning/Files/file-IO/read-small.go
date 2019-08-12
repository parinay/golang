package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func readSmall() {
	fptr := flag.String("fpath1", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	buffer := make([]byte, 3)

	for {
		n, err := r.Read(buffer)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(buffer[0:n]))
	}
}
