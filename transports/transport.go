package transports

import (
	"context"
	"errors"
	demo "github.com/go-kit-demo/demo_pb"
	endpoints "github.com/go-kit-demo/endpopints"
	"github.com/go-kit-demo/entities"
	"github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	demo.UnimplementedGreeterServer
	grpc1 grpc.Handler
}

func NewgRPCServer(endpt endpoints.Endpoints) demo.GreeterServer {
	return &gRPCServer{
		grpc1: grpc.NewServer(
			endpt.SayHello,
			decodeGreetingRequest,
			encodeGreetingResponse),
	}
}

func (g *gRPCServer) SayHello(ctx context.Context, req *demo.HelloRequest) (*demo.HelloResponse, error) {
	_, res, err := g.grpc1.ServeGRPC(ctx, req)
	return res.(*demo.HelloResponse), err
}

func decodeGreetingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*demo.HelloRequest)
	if !ok {
		return nil, errors.New("invalid request body")
	}

	return entities.HelloRequest{Name: req.Name}, nil
}

func encodeGreetingResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(entities.HelloResponse)
	if !ok {
		return nil, errors.New("invalid response body")
	}

	return &demo.HelloResponse{Greeting: resp.Greeting}, nil
}
