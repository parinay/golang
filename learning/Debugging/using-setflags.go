package main

import "log"

func init() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func example() {

	// Will print: "[date] [time] [filename]:[line]: [text]"
	log.Println("Logging w/ line numbers on golangcode.com")
}

func call1() {
	log.Println("Logging w/ line numbers on golangcode.com")
}
func usingSetFlags() {
	log.Println("Logging w/ line numbers on golangcode.com")
	call1()
	example()
}
