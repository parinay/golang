package main

import (
	"fmt"
	"path/filepath"
)

func assertCompare() {
	files, error := filepath.Glob("[")

	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println("error: ", error)
		return
	}
	fmt.Println("Matched files ", files)
}
