// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"

	"iterative_control/internal/config"
	"iterative_control/internal/handler"
	"iterative_control/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/iterative.yaml", "the config file")

func main() {

	logx.Debug("Debug level log")
	logx.Info("Info level log")
	logx.Error("Error level log")

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//打印配置信息
	fmt.Printf("Config: %#v\n", c)

	//server := rest.MustNewServer(c.RestConf) //不开启跨域
	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
