## 依赖安装
### 安装 Protocol Buffers（protoc）
官网：https://github.com/protocolbuffers/protobuf/releases      可以点击 Tags 切换版本，最新版本的可能找不到 windows 版本

比如下载：protoc-33.0-win64.zip，然后解压，把解压后的bin目录添加到 环境变量即可。

为了方便管理，建议解压到 GOPATH 下面
```shell
go env GOPATH
```

### 安装 Go 的 gRPC 相关工具
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 依赖会安装到 %GOPATH%\bin 目录下
where protoc-gen-go
D:\soft\go_path\bin\protoc-gen-go.exe

where protoc-gen-go-grpc
D:\soft\go_path\bin\protoc-gen-go-grpc.exe
```

## Demo
创建 grpc/hello_grpc 目录，整体的目录结果如下：
```
hello_grpc
    ├─client
    │      main.go
    │
    ├─proto
    │      hello_grpc.pb.go
    │      hello_grpc.proto
    │      hello_grpc_grpc.pb.go
    │
    └─server
            main.go
```

在 proto 目录下执行：protoc --go_out=. --go-grpc_out=. hello_grpc.proto 即可



## 多 proto 文件
原则：每个 proto 一个 Go 包
```
proto/
├── basic/
│   └── basic.proto
├── order/
│   └── order.proto
└── user/
    └── user.proto
```

protoc 
--go_out=module=go_learning/04_micro/grpc/multiple_proto/proto:. 
--go-grpc_out=module=go_learning/04_micro/grpc/multiple_protogo_learning/04_micro/grpc/multiple_proto/proto:. 
basic/basic.proto 
order/order.proto 
user/user.proto

protoc --go_out=module=go_learning/04_micro/grpc/multiple_proto/proto:. --go-grpc_out=module=go_learning/04_micro/grpc/multiple_protogo_learning/04_micro/grpc/multiple_proto/proto:. basic/basic.proto order/order.proto user/user.proto
