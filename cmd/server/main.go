package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rivory/gogrpcvshttp/handler"
	service "github.com/rivory/gogrpcvshttp/pkg/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type serverHTTP struct {
	handler handler.Hhttp
}

type serverGRPC handler.Hgrpc

func main() {
	serverHTTP := serverHTTP{
		handler: handler.ProvideHTTPHandler(),
	}
	http.HandleFunc("/hello", serverHTTP.handler.Handle)

	go func() {
		fmt.Printf("Starting http server at port 8080\n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Print("Grpc Server listening")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterServiceServer(s, handler.ProvideGrpcHandler())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
