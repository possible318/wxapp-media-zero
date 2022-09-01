package logic

import (
	"context"

	"media-zero/service/photos/rpc/internal/svc"
	"media-zero/service/photos/rpc/photo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PicsLogic {
	return &PicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PicsLogic) Pics(in *photo.PicsReq) (*photo.PicsRes, error) {
	// todo: add your logic here and delete this line
	// 获取参数
	id := in.Id
	// 从 blog 中获取数据
	res, err := l.svcCtx.BlogModel.FindOne(l.ctx, uint64(id))
	if err != nil {
		return nil, err
	}

	return &photo.PicsRes{
		Id:     int64(res.Id),
		ItemId: res.ItemId,
		Pid:    res.Pid,
		Text:   res.Pid,
		Src:    res.Src,
	}, nil
}
