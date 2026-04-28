package grpc

import (
	"auth/internal/svc"
	"auth/internal/types"
	"context"
	"strconv"

	"backend/common/protocol/authpb"
)

type AuthGRPCServer struct {
	authpb.UnimplementedAuthServiceServer
	svcCtx *svc.ServiceContext
}

func NewAuthGRPCServer(svcCtx *svc.ServiceContext) *AuthGRPCServer {
	return &AuthGRPCServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthGRPCServer) VerifyToken(ctx context.Context, req *authpb.VerifyTokenRequest) (*authpb.VerifyTokenResponse, error) {
	claims, err := s.svcCtx.JWTMgr.ValidateToken(req.AccessToken, false)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseInt(claims.UserID, 10, 64)
	if err != nil {
		return nil, err
	}

	savedToken, err := s.svcCtx.TokenDAO.GetAccessToken(ctx, userID)
	if err != nil {
		return nil, err
	}
	if savedToken == "" {
		return nil, types.ErrTokenRevoked
	}
	if savedToken != req.AccessToken {
		return nil, types.ErrTokenRevoked
	}

	return &authpb.VerifyTokenResponse{
		UserId:   userID,
		Role:     claims.Role,
		Username: claims.Username,
	}, nil
}
