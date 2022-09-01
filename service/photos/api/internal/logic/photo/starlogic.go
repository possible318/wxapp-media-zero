package photo

import (
	"context"

	"media-zero/service/photos/api/internal/svc"
	"media-zero/service/photos/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarLogic {
	return &StarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StarLogic) Star(req *types.StarReq) (resp *types.StarRes, err error) {
	id := req.ID
	res, err := l.svcCtx.BlogModel.FindOne(l.ctx, uint64(id))
	if err != nil {
		return nil, err
	}
	// 点赞数加1
	res.Like++
	if err := l.svcCtx.BlogModel.Update(l.ctx, res); err != nil {
		return nil, err
	}
	return &types.StarRes{}, nil
}
