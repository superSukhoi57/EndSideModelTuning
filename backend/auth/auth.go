package main

import (
	"flag"
	"fmt"
	"net"

	"auth/internal/config"
	authgrpc "auth/internal/grpc"
	"auth/internal/handler"
	"auth/internal/svc"
	"backend/common/protocol/authpb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	restServer := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	handler.RegisterHandlers(restServer, svcCtx)
	serviceGroup.Add(restServer)
	fmt.Printf("Starting rest server at %s:%d...\n", c.Host, c.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.GRPC.Port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen grpc port: %v", err))
	}

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, authgrpc.NewAuthGRPCServer(svcCtx))
	reflection.Register(grpcServer)
	serviceGroup.Add(service.WithStart(func() {
		fmt.Printf("Starting grpc server at :%d...\n", c.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			panic(fmt.Sprintf("failed to serve grpc: %v", err))
		}
	}))

	serviceGroup.Start()
}
