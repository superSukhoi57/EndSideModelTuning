package middleware

import (
	"context"
	"net/http"
	"strings"

	"backend/common/protocol/authpb"
	"iterative_control/internal/svc"
)

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
	RoleKey   ContextKey = "role"
)

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

		resp, err := m.svcCtx.AuthClient.VerifyToken(r.Context(), &authpb.VerifyTokenRequest{
			AccessToken: token,
		})
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, resp.GetUserId())
		ctx = context.WithValue(ctx, RoleKey, resp.GetRole())

		next(w, r.WithContext(ctx))
	}
}
