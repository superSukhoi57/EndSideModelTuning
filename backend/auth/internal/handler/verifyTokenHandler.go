package handler

import (
	"auth/internal/logic"
	"auth/internal/svc"
	"auth/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func verifyTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVerifyTokenLogic(r.Context(), svcCtx)
		resp, err := l.VerifyToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
