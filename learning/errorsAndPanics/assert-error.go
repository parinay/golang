package main

import (
	"fmt"
	"os"
)

func assertError() {

	f, err := os.Open("./test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("err: File at path", err.Path, "failed to open")
		fmt.Println("err", err)
		return
	}

	fmt.Println(f.Name(), "opened successfully.")
}
