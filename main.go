package main

import (
	demo "github.com/go-kit-demo/demo_pb"
	endpoints "github.com/go-kit-demo/endpopints"
	"github.com/go-kit-demo/services"
	"github.com/go-kit-demo/transports"
	"github.com/go-kit/log"
	"google.golang.org/grpc"

	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()

	logger := log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	s := services.NewService(logger)
	endp := endpoints.MakeEndpoints(s)
	g := transports.NewgRPCServer(endp)

	demo.RegisterGreeterServer(server, g)
	logger.Log("gRPC server is running on port 3000")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}

}
