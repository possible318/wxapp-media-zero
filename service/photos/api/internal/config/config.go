package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf
	Mysql      struct {
		DataSource string
	}
	PicsRpc zrpc.RpcClientConf
}
