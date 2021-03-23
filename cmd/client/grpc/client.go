package main

import (
	"context"
	"fmt"
	"log"
	"time"

	service "github.com/rivory/gogrpcvshttp/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := service.NewServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	start := time.Now()
	r, err := c.Handle(ctx, &service.HelloWorld{Message: "tototest"})
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("http.Post took %v microseconds \n", elapsed)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
