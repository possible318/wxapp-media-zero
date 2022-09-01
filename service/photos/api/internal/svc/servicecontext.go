package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"media-zero/service/photos/api/internal/config"
	"media-zero/service/photos/model"
)

type ServiceContext struct {
	Config    config.Config
	BlogModel model.BlogModel
	//PicsRpc   pics.Pics
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		BlogModel: model.NewBlogModel(conn, c.CacheRedis), // 初始化 BlogModel
		//PicsRpc:   pics.NewPics(zrpc.MustNewClient(c.PicsRpc)), // 初始化 UserRpc
	}
}
