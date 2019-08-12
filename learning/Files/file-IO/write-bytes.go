package main

import (
	"fmt"
	"os"
)

func writeBytes() {
	fd, err := os.Create("bytes.txt")

	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}

	bytes := []byte{104, 101, 111, 33, 44, 181, 191, 255}

	count, err := fd.Write(bytes)

	if err != nil {
		fmt.Println("Failed to write", err)
		fd.Close()
		return
	}

	fmt.Println(count, "bytes written successfully")

	err = fd.Close()
	if err != nil {
		fmt.Println("Failed to close successfully ", err)
		return
	}
}
