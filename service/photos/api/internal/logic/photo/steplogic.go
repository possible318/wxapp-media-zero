package photo

import (
	"context"

	"media-zero/service/photos/api/internal/svc"
	"media-zero/service/photos/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StepLogic {
	return &StepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StepLogic) Step(req *types.StarReq) (resp *types.StarRes, err error) {
	id := req.ID
	res, err := l.svcCtx.BlogModel.FindOne(l.ctx, uint64(id))
	if err != nil {
		return nil, err
	}
	// 点赞数加1
	res.Dislike++
	if err := l.svcCtx.BlogModel.Update(l.ctx, res); err != nil {
		return nil, err
	}
	return &types.StarRes{}, nil
}
