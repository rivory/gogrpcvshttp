package handler

import (
	"context"
	"log"

	service "github.com/rivory/gogrpcvshttp/pkg/proto"

	"github.com/rivory/gogrpcvshttp/domain"
)

// Hgrpc composes the grpc handler
type Hgrpc struct {
	service domain.LogicInterface
	service.UnimplementedServiceServer
}

// ProvideGrpcHandler provides grpc handler implementation
func ProvideGrpcHandler() *Hgrpc {
	return &Hgrpc{
		service: domain.LogicService{},
	}
}

// Handle handles grpc request
func (handler *Hgrpc) Handle(ctx context.Context, in *service.HelloWorld) (*service.HelloWorld, error) {
	log.Printf("grpc Received: %v", in.Message)

	in.Message = handler.service.Uppercase(in.Message)

	return in, nil
}
