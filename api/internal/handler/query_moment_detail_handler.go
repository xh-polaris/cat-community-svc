package handler

import (
	"net/http"

	"github.com/xh-polaris/cat-community-svc/api/internal/logic"
	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryMomentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryMomentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQueryMomentDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryMomentDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
