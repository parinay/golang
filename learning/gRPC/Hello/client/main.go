package main

import (
	"context"
	"log"

	pb "github.com/parinay/golang/learning/gRPC/Hello/sample"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	r, err = c.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
