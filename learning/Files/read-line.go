package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func readLine() {
	fptr := flag.String("falth2", "test.txt", "file path to read from")
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

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
