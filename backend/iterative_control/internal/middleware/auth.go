package middleware

import (
	"context"
	"net/http"
	"strings"

	"backend/common/enumeration"
	"backend/common/protocol/authpb"
	"iterative_control/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContextKey string

type AuthMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewAuthMiddleware(svcCtx *svc.ServiceContext) *AuthMiddleware {
	return &AuthMiddleware{svcCtx: svcCtx}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		logx.Infof("请求头的认证信息: %s", authHeader)

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			token = authHeader
		}

		resp, err := m.svcCtx.AuthClient.VerifyToken(r.Context(), &authpb.VerifyTokenRequest{
			AccessToken: token,
		})
		logx.Infof("grpc返回: %v", resp)
		if err != nil {
			logx.Errorf("rpc调用出错===== 》error: %v", err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		logx.Infof("令牌通过校验-----> user_id: %d, role: %s", resp.GetUserId(), resp.GetRole())
		ctx := context.WithValue(r.Context(), enumeration.UserIDKey, resp.GetUserId())
		ctx = context.WithValue(ctx, enumeration.RoleKey, resp.GetRole())

		next(w, r.WithContext(ctx))
	}
}
