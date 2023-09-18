package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mikekbnv/gRPC-echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type EchoServer struct {
	echo.EchoServer
}

func (s *EchoServer) Echo(ctx context.Context, n *echo.Msg) (*echo.Reply, error) {
	log.Printf("Recieved a msg: %v", n.Text)

	text, err := echo.Echo(n)
	if err != nil {
		return &echo.Reply{}, err
	}

	return text, nil
}
func main() {
	// parse arguments from the command line
	// this lets us define the port for the server
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// Check for errors
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Instantiate the server
	s := grpc.NewServer()
	reflection.Register(s)
	echo.RegisterEchoServer(s, &EchoServer{})
	// Register server method (actions the server will do)
	// TODO

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
