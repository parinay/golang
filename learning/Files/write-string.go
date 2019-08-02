package main

import (
	"fmt"
	"os"
)

func writeString() {
	fd, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	count, err := fd.WriteString("Hello World!")

	if err != nil {
		fmt.Println(err)
		fd.Close()
		return
	}

	fmt.Println(count, "bytes written successfully")

	err = fd.Close()

	if err != nil {
		fmt.Println("Problems closing files", err)
		return
	}
}
