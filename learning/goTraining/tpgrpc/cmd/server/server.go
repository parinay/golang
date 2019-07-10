package main

import (
	"context"
	"net"
	"time"

	"github.com/parinay/golang/goTraining/tpgrpc/pkg/tpgrpc"
	"google.golang.org/grpc"
)

type ATpServer struct {
}

func (*ATpServer) Hello(ctx context.Context, req *tpgrpc.TPRequest) (*tpgrpc.TPResponse, error) {
	return &tpgrpc.TPResponse{Text: "Hi"}, nil
}
func (*ATpServer) Time(ctx context.Context, req *tpgrpc.TPRequest) (*tpgrpc.TPResponse, error) {
	return &tpgrpc.TPResponse{Text: time.Now().String()}, nil
}
func (*ATpServer) Greetings(req *tpgrpc.TPRequest, srv tpgrpc.Tp_GreetingsServer) error {
	srv.Send(&tpgrpc.TPResponse{Text: "Good morning"})
	srv.Send(&tpgrpc.TPResponse{Text: "Good env"})
	srv.Send(&tpgrpc.TPResponse{Text: "Good night!"})
	return nil
}
func main() {
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		return
	}
	srv := grpc.NewServer()
	tpgrpc.RegisterTpServer(srv, &ATpServer{})
	srv.Serve(l)
}
