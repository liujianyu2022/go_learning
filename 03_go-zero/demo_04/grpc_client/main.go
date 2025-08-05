package main

import (
	"context"
	"log"

	"go_learning/03_go-zero/demo_04/grpc_server/greet"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/config.yaml", &clientConf)

	// 这是使用 go zero 框架封装的方式，这是 Go-Zero 对原生 gRPC 客户端的封装，简化了配置和连接管理，集成 Go-Zero 的服务发现、负载均衡等能力。
	// conn := zrpc.MustNewClient(clientConf)				
	// client := greet.NewGreetClient(conn.Conn())

	// 这是使用 grpc 原生方式，创建一个 gRPC 客户端连接
	conn, err := grpc.NewClient(						
		"etcd://127.0.0.1:2379/greet_rpc", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := greet.NewGreetClient(conn)
	
	response, err := client.Ping(context.Background(), &greet.Request{Ping: "ping"})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("response = ", response)
}
