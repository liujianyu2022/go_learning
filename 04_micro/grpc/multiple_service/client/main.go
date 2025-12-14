package main

import (
	"context"
	mutiple_service "go_learning/04_micro/grpc/multiple_service/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	orderClient := mutiple_service.NewOrderServiceClient(conn)
	videoClient := mutiple_service.NewVideoServiceClient(conn)

	orderResult, err := orderClient.GetOrder(
		context.Background(),
		&mutiple_service.Request{Name: "Test Order"},
	)
	if err != nil {
		log.Fatalf("could not get order: %v", err)
	}
	log.Printf("Order Service Response: %s", orderResult.Name)

	videoResult, err := videoClient.GetVideo(
		context.Background(),
		&mutiple_service.Request{Name: "Test Video"},
	)
	if err != nil {
		log.Fatalf("could not get video: %v", err)
	}
	log.Printf("Video Service Response: %s", videoResult.Name)
}