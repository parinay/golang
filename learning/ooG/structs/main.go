package main

import "github.com/parinay/golang/learning/ooG/structs/employee"

func main() {
	e := employee.New("sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
