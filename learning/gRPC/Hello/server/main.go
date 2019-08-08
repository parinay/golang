package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/parinay/golang/learning/gRPC/Hello/sample"
	"google.golang.org/grpc"
)

type server struct {
	helloRequest *pb.HelloRequest
}

const (
	port = ":50052"
)

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
	fmt.Println("vim-go")
}
