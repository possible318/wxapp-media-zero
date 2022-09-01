package photo

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"media-zero/service/photos/api/internal/svc"
	"media-zero/service/photos/api/internal/types"
)

type PicsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PicsLogic {
	return &PicsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PicsLogic) Pics(req types.PicsReq) (resp *types.PicsRes, err error) {
	// 获取参数
	page := req.Page
	// 从 blog 中获取数据
	res, err := l.svcCtx.BlogModel.PageList(l.ctx, page, 10)
	if err != nil {
		return nil, err
	}
	// 获取total
	count, err := l.svcCtx.BlogModel.Count(l.ctx)
	if err != nil {
		return nil, err
	}

	picList := make([]types.PicItem, 0)
	for _, item := range res {
		picItem := types.PicItem{
			ID:     int(item.Id),
			ItemId: item.ItemId,
			Pid:    item.Pid,
			Text:   item.Pid,
			Src:    item.Src,
		}
		picList = append(picList, picItem)
	}

	data := &types.PicList{
		Total: count,
		List:  picList,
	}

	return &types.PicsRes{
		Code: 0,
		Msg:  "success",
		Data: *data,
	}, nil
}
