package handler

import (
	"auth/internal/logic"
	"auth/internal/svc"
	"auth/internal/types"
	"fmt"
	"net/http"
	"net/url"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func callbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCallbackLogic(r.Context(), svcCtx)
		resp, err := l.Callback(&req)
		if err != nil {
			logx.Errorf("Callback error: %v", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		frontendURL := svcCtx.Config.Frontend.RedirectURI
		params := url.Values{}
		params.Set("access_token", resp.AccessToken)
		params.Set("refresh_token", resp.RefreshToken)
		params.Set("user_info", fmt.Sprintf(
			`{"open_id":"%s","union_id":"%s","name":"%s","en_name":"%s","email":"%s","mobile":"%s","avatar_url":"%s","tenant_name":"%s"}`,
			resp.UserInfo.OpenId,
			resp.UserInfo.UnionId,
			resp.UserInfo.Name,
			resp.UserInfo.EnName,
			resp.UserInfo.Email,
			resp.UserInfo.Mobile,
			resp.UserInfo.AvatarUrl,
			resp.UserInfo.TenantName,
		))

		redirectURL := fmt.Sprintf("%s?%s", frontendURL, params.Encode())
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}
