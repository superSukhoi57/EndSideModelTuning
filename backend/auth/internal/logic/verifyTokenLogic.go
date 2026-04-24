package logic

import (
	"auth/internal/svc"
	"auth/internal/types"
	"context"
	"strconv"
	"time"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyTokenLogic) VerifyToken(req *types.VerifyTokenReq) (*types.VerifyTokenResp, error) {
	claims, err := l.svcCtx.JWTMgr.ValidateToken(req.AccessToken, false)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.ParseInt(claims.UserID, 10, 64)
	if err != nil {
		return nil, err
	}

	savedToken, err := l.svcCtx.TokenDAO.GetAccessToken(l.ctx, userID)
	if err != nil {
		return nil, err
	}
	if savedToken == "" {
		return nil, types.ErrTokenRevoked
	}
	if savedToken != req.AccessToken {
		return nil, types.ErrTokenRevoked
	}

	return &types.VerifyTokenResp{
		Base: types.BaseResp{
			Code:      0,
			Message:   "success",
			Timestamp: time.Now().Unix(),
		},
		UserID: userID,
		Role:   claims.Role,
	}, nil
}
