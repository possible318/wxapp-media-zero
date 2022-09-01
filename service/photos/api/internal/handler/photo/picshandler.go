package photo

import (
	"media-zero/service/photos/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"media-zero/service/photos/api/internal/logic/photo"
	"media-zero/service/photos/api/internal/svc"
)

func PicsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PicsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := photo.NewPicsLogic(r.Context(), svcCtx)
		resp, err := l.Pics(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
