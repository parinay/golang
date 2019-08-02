package main

import (
	"fmt"
	"os"
)

func writeLine() {
	fd, err := os.Create("lines.txt")
	if err != nil {
		fmt.Println("Failed to create file", err)
		fd.Close()
		return
	}

	lines := []string{"Welcome to the world of Go!", "Go is a compiled language", "Its easy to learn go", "Practice, practice, practice!"}

	for _, v := range lines {
		_, err := fmt.Fprintln(fd, v)
		if err != nil {
			fmt.Println("Failed to write line in file", err)
			fd.Close()
			return
		}
	}

	// fmt.Println(count, "lines written successfully")
	err = fd.Close()
	if err != nil {
		fmt.Println("Failed to close file", err)
		return
	}

	fmt.Println("File written successfully")
}
