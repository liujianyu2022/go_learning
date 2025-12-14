package main

import (
	"context"
	hello_grpc "go_learning/04_micro/grpc/hello_grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := hello_grpc.NewHelloServiceClient(conn)
	result, err := client.SayHello(
		context.Background(), 
		&hello_grpc.HelloRequest{Name: "gRPC Client"},
	)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s, Message: %s", result.Name, result.Message)
}