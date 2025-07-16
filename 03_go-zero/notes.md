## go-zero 快速入门

### 安装与环境配置
1. goctl
goctl 是 go-zero框架 的内置脚手架，可以一键生成代码、文档、部署k8s yaml、dockerfile等
```
go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl --version
```
go get 仅下载代码并更新 go.mod，不再安装二进制文件。
go install 会直接将二进制文件安装到 GOPATH/bin。安装二进制工具需要使用 go install

2. protoc
- protobuf 是一种序列化/反序列化的标准，类似于json、xml
- protoc命令行是protobuf的编译器，将 .proto 文件编译成对应的开发语言文件
- protoc-gen-go 是protoc的一个插件，用于生成go语言代码
- protoc-gen-go-grpc是protoc的go grpc的插件，可以生成grpc相关的go语言文件

打开https://github.com/protocolbuffers/protobuf/releases
下载对应的版本（我这里是windows电脑），下载 https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protoc-28.0-win64.zip
解压，并设置环境变量即可  protoc --version

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

3. IDEA插件
- goctl


### 创建 demo
1. 创建步骤
```
cd 03_go-zero               
goctl api new demo_01       // 使用脚手架创建模板代码
cd demo_01
go mod tidy                 // 下载依赖
```

2. 目录说明
```
example                             // example：单个服务目录，一般是某微服务名称
├── etc                             // etc：静态配置文件目录
│   └── example.yaml
├── main.go                         // main.go：程序启动入口文件
└── internal                        // internal：单个服务内部文件，其可见范围仅限当前服务
    ├── config                      // config：静态配置文件对应的结构体声明目录
    │   └── config.go
    ├── handler                     // handler：handler 目录，可选，一般 http 服务会有这一层做路由管理，handler 为固定后缀
    │   ├── xxxhandler.go
    │   └── xxxhandler.go
    ├── logic                       // logic：业务目录，所有业务编码文件都存放在这个目录下面，logic 为固定后缀
    │   └── xxxlogic.go
    ├── svc                         // svc：依赖注入目录，所有 logic 层需要用到的依赖都要在这里进行显式注入
    │   └── servicecontext.go
    └── types                       // types：结构体存放目录
        └── types.go
```

3. 修改 demo_01.api
不要直接修改 types 目录下的内容，这是 goctl 自动生成的
修改 demo_01.api 之后，重新生成文件

goctl api go --api demo_01.api --dir .

--api：指定api文件
--dir：指定go文件生成的目录


### 配置文件
在生成的代码中，配置文件位于etc目录下，格式为yaml。go-zero支持多种配置文件格式：yml、yaml、toml、json


在部署时，**配置文件并不会打包到二进制包中**，我们需要指定配置文件，通过命令行参数的形式指定
```
# 执行文件 -f etc/hello01-api.json
go run main.go -f etc/hello01-api.json
```