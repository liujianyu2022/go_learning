package main

import (
	"context"
	hello_grpc "go_learning/04_micro/grpc/hello_grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServer struct {
	hello_grpc.UnimplementedHelloServiceServer
}

func (*HelloServer) SayHello(
	ctx context.Context,
	req *hello_grpc.HelloRequest,
) (*hello_grpc.HelloResponse, error) {

	return &hello_grpc.HelloResponse{
		Name:    "Hello " + req.Name,
		Message: "This is a gRPC Hello World message: " + req.Message,
	}, nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器实例
	grpcServer := grpc.NewServer()

	// 注册服务
	hello_grpc.RegisterHelloServiceServer(grpcServer, &HelloServer{})

	// 启动前打印日志
	log.Println("gRPC server is running on port 50051")

	// 启动服务器（阻塞）
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
