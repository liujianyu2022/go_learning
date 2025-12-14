package main

import (
	"context"
	"log"
	"net"

	multiple_service "go_learning/04_micro/grpc/multiple_service/proto"

	"google.golang.org/grpc"
)

type VideoService struct{
	multiple_service.UnimplementedVideoServiceServer
}
type OrderService struct{
	multiple_service.UnimplementedOrderServiceServer
}

func (*VideoService) GetVideo(ctx context.Context, req *multiple_service.Request) (*multiple_service.Response, error) {
	return &multiple_service.Response{
		Name: req.Name + " Video Service Response",
	}, nil
}

func(*OrderService) GetOrder(ctx context.Context, req *multiple_service.Request) (*multiple_service.Response, error) {
	return &multiple_service.Response{
		Name: req.Name + " Order Service Response",
	}, nil
}

func main(){
	// 监听端口
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器实例
	grpcServer := grpc.NewServer()

	// 注册服务
	multiple_service.RegisterVideoServiceServer(grpcServer, &VideoService{})
	multiple_service.RegisterOrderServiceServer(grpcServer, &OrderService{})

	log.Println("gRPC Multiple Service server is running on port 50052")

	// 启动grpc服务器（阻塞）
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}