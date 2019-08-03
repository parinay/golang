package main

import (
	"fmt"
	"os"
)

func appendFile() {
	fd, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("File open in append mode failed:", err)
		return
	}
	newLine := "File handling is easy"

	_, err = fmt.Fprintln(fd, newLine)
	if err != nil {
		fd.Close()
		fmt.Println("Write in append mode failed", err)
		return
	}

	err = fd.Close()
	if err != nil {
		fmt.Println("File close failed:", err)
		return
	}

	fmt.Println("file appended successfully")
}
