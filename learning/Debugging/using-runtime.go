package main

import (
	"fmt"
	"runtime"
	"strings"
)

func fileName(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

func debugLog(msg string) {

	programCounter, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)

	prefix := fmt.Sprintf("[%s:%s():%d]: %s", fileName(file), fn.Name(), line, msg)

	fmt.Printf(prefix)
	fmt.Println()
}

func exampleR() {
	debugLog("testing the msg")
}
func usingRuntime() {
	exampleR()
}
