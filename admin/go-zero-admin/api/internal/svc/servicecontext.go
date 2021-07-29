// logic所依赖的资源池

package svc

import (
	"go-zero-admin/api/internal/config"
	"go-zero-admin/api/internal/middleware"
	"go-zero-admin/rpc/oms/omsclient"
	"go-zero-admin/rpc/pms/pmsclient"
	"go-zero-admin/rpc/sms/smsclient"
	"go-zero-admin/rpc/sys/sysclient"
	"go-zero-admin/rpc/ums/umsclient"

	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config   //配置声明type
	CheckUrl rest.Middleware //中间件
	Sys      sysclient.Sys
	Ums      umsclient.Ums
	Pms      pmsclient.Pms
	Oms      omsclient.Oms
	Sms      smsclient.Sms
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.NewRedis(c.Redis.Address, redis.NodeType)
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		Sys:      sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums:      umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms:      pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms:      omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms:      smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Redis:    newRedis,
	}
}