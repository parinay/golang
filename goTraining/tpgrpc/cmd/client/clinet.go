package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	tpgrpc "github.com/parinay/golang/goTraining/tpgrpc/pkg/tpgrpc"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	client := tpgrpc.NewTpClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Hello(ctx, &tpgrpc.TPRequest{})
	if err != nil {
		log.Println("Error from server for Hello() :", err)
		return
	}
	fmt.Printf("Server Response  - %s\n", resp.Text)

	resp, err = client.Time(ctx, &tpgrpc.TPRequest{})
	if err != nil {
		log.Println("Error from server for Time() :", err)
		return
	}
	fmt.Printf("Server Response  - %s\n", resp.Text)

	gc, err := client.Greetings(ctx, &tpgrpc.TPRequest{})
	if err != nil {
		log.Println("Error from Server for Greetings()", err)
		return
	}
	for {
		resp, err = gc.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error ", err)
			return
		}
		fmt.Printf("Response : %s\n", resp.Text)
	}

}
