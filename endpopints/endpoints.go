package endpoints

import (
	"context"
	"github.com/go-kit-demo/entities"
	"github.com/go-kit-demo/services"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct holds a list of endpoints
type Endpoints struct {
	SayHello endpoint.Endpoint
	//Add      endpoint.Endpoint
	//Subtract endpoint.Endpoint
	//Multiply endpoint.Endpoint
	//Divide   endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		SayHello: makesayHelllo(s),
	}
}

func makesayHelllo(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entities.HelloRequest)
		return s.SayHello(ctx, req.Name), nil
	}
}
