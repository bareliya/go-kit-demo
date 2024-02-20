package services

import (
	"context"
	"github.com/go-kit/log"
)

type service struct {
	logger log.Logger
}

// Service interface describes a service that adds numbers
type Service interface {
	SayHello(ctx context.Context, name string) string
	//Add(ctx context.Context, numA, numB float32) (float32, error)
	//Subtract(ctx context.Context, numA, numB float32) (float32, error)
	//Multiply(ctx context.Context, numA, numB float32) (float32, error)
	//Divide(ctx context.Context, numA, numB float32) (float32, error)
}

// NewService creates a new service with all the expected dependencies and returns it
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) SayHello(ctx context.Context, name string) string {
	return "Hello" + name
}
