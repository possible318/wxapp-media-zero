package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"media-zero/service/photos/model"
	"media-zero/service/photos/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	BlogModel model.BlogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		BlogModel: model.NewBlogModel(conn, c.CacheRedis), // 初始化 BlogModel
	}
}
