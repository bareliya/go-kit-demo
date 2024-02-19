package main

import (
	"context"
	"fmt"
	demo "github.com/go-kit-demo/demo_pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	var g GreeterServer
	demo.RegisterGreeterServer(server, &g)
	log.Println("gRPC server is running on port 3000")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	fmt.Println("hello world")
}

type GreeterServer struct {
}

func (g *GreeterServer) SayHello(ctx context.Context, req *demo.HelloRequest) (*demo.HelloResponse, error) {
	return &demo.HelloResponse{Greeting: "hello" + req.Name}, nil
}

//func (g *GreeterServer) mustEmbedUnimplementedGreeterServer() {
//	// do nothing
//}
