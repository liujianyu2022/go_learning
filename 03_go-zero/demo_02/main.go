package main

import (
	"flag"
	"fmt"

	"go_learning/03_go-zero/demo_02/internal/config"
	"go_learning/03_go-zero/demo_02/internal/handler"
	"go_learning/03_go-zero/demo_02/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)				// 服务上下文 依赖注入，需要用到的依赖都在此进行注入，比如配置，数据库连接，redis连接等
	handler.RegisterHandlers(server, ctx)		// 注册路由

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
