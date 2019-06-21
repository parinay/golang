package main

import (
	"fmt"
	"log"
	"net"
)

type tpclient struct {
	conn net.Conn
}

func (c *tpclient) hello() string {
	var output string
	fmt.Fprintln(c.conn, "hello")
	fmt.Fscanf(c.conn, "%s\n", &output)

	return output
}

func NewTPclient(c net.Conn) *tpclient {
	return &tpclient{
		conn: c,
	}
}
func main() {
	cln, err := net.Dial("tcp", "localhost:8087")
	if err != nil {
		log.Printf("Connection error:%v", err)
	}
	c := NewTPclient(cln)
	fmt.Printf("Server response: %v\n", c.hello())

	var output string
	fmt.Fprintln(cln, "time")
	fmt.Fscanf(cln, "%s\n", &output)
	fmt.Printf("Server response: %v\n", output)

	fmt.Fprintln(cln, "close")

}
