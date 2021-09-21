package main

import (
	"flag"
	"fmt"
	"go-zero-admin/api/internal/common/errorx"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"

	"go-zero-admin/api/internal/config"
	"go-zero-admin/api/internal/handler"
	"go-zero-admin/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "api/etc/admin-api.yaml", "the config file")

//var configFile = flag.String("f", "./etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) //加载配置文件

	ctx := svc.NewServiceContext(c)          //logic所依赖的资源池
	server := rest.MustNewServer(c.RestConf) //创建Rest服务
	defer server.Stop()                      //停止服务

	handler.RegisterHandlers(server, ctx) //注册Handler及路由設置

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start() //启动服务
}
