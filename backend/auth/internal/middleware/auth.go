package middleware

import (
	"auth/internal/svc"
	"backend/common/enumeration"
	"context"
	"net/http"
	"strconv"
	"strings"

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

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			token = authHeader
		}

		claims, err := m.svcCtx.JWTMgr.ValidateToken(token, false)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		userID, err := strconv.ParseInt(claims.UserID, 10, 64)
		if err != nil {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		savedToken, err := m.svcCtx.TokenDAO.GetAccessToken(r.Context(), userID)
		if err != nil {
			http.Error(w, "failed to verify token", http.StatusInternalServerError)
			return
		}
		if savedToken == "" || savedToken != token {
			http.Error(w, "token revoked", http.StatusUnauthorized)
			return
		}
		logx.Infof("令牌通过校验-----> user_id: %d, role: %s", userID, claims.Role)

		ctx := context.WithValue(r.Context(), enumeration.UserIDKey, userID)
		ctx = context.WithValue(ctx, enumeration.RoleKey, claims.Role)

		next(w, r.WithContext(ctx))
	}
}
