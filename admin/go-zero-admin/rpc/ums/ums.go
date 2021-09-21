// Code generated by goctl. DO NOT EDIT!
// Source: ums.proto

package main

import (
	"flag"
	"fmt"

	"go-zero-admin/rpc/ums/internal/config"
	"go-zero-admin/rpc/ums/internal/server"
	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "rpc/ums/etc/ums.yaml", "the config file")

//var configFile = flag.String("f", "./etc/ums.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUmsServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ums.RegisterUmsServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
