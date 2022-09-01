package photo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"media-zero/service/photos/api/internal/logic/photo"
	"media-zero/service/photos/api/internal/svc"
	"media-zero/service/photos/api/internal/types"
)

func StepHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StarReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := photo.NewStepLogic(r.Context(), svcCtx)
		resp, err := l.Step(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
