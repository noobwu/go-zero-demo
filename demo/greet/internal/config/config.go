//配置声明type

package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf                 // rest api配置
	CacheRedis    cache.CacheConf // redis缓存
	Auth          struct {        // jwt鉴权配置

		AccessSecret string
		AccessExpire int64
	}
}
