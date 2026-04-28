package handler

import (
	"auth/internal/logic"
	"auth/internal/svc"
	"auth/internal/types"
	"net/http"

	"backend/common/enumeration"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func refreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		userID := r.Context().Value(enumeration.UserIDKey).(int64)

		l := logic.NewRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.RefreshToken(&req, userID)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
