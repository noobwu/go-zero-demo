package main

import (
	"flag"
	"fmt"

	"go-zero-admin/front-api/internal/config"
	"go-zero-admin/front-api/internal/handler"
	"go-zero-admin/front-api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "front-api/etc/front-api.yaml", "the config file")

//var configFile = flag.String("f", "./etc/front-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
