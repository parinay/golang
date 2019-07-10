package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleconnection(conn net.Conn) {
	log.Println("Connection started.")
	var command string
OuterLoop:
	for {
		_, err := fmt.Fscanf(conn, "%s\n", &command)
		if err != nil {
			log.Printf("Connection error:%v", err)
			_, ok := err.(net.Error)
			if ok {
				break
			}
		}
		switch command {
		case "hello":
			fmt.Fprintln(conn, "hi")
		case "time":
			fmt.Fprintln(conn, time.Now())
		case "close":
			conn.Close()
			break OuterLoop
		default:
			fmt.Fprintln(conn, "Inavlid use of TP protocol")
		}
	}
	log.Println("connection closed.")
}
func main() {
	log.Printf("Starting the server ...")
	srv, err := net.Listen("tcp", ":8087")
	if err != nil {
		//handle error
		log.Printf("Failed to start server. Error:%v", err)
		return
	}
	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Printf("Failed to accept connection req. Error: %v\n", err)
			return
		}
		go handleconnection(conn)

	}
	log.Printf("Stopping the server ...")
}
