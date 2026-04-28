package handler

import (
	"auth/internal/logic"
	"auth/internal/svc"
	"auth/internal/types"
	"net/http"

	"backend/common/enumeration"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(enumeration.UserIDKey).(int64)

		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(&types.LogoutReq{UserID: userID})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
