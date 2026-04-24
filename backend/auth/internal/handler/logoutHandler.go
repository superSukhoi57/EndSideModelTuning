package handler

import (
	"auth/internal/logic"
	"auth/internal/svc"
	"auth/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogoutReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
